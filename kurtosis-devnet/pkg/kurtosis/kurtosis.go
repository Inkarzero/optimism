package kurtosis

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

const DefaultPackageName = "github.com/ethpandaops/optimism-package"

type EndpointMap map[string]string

type Chain struct {
	Name      string              `json:"name"`
	ID        string              `json:"id,omitempty"`
	Endpoints EndpointMap         `json:"endpoints"`
	Addresses DeploymentAddresses `json:"addresses,omitempty"`
}

// KurtosisEnvironment represents the output of a Kurtosis deployment
type KurtosisEnvironment struct {
	L1 *Chain   `json:"l1"`
	L2 []*Chain `json:"l2"`
}

// KurtosisDeployer handles deploying packages using Kurtosis
type KurtosisDeployer struct {
	// Base directory where the deployment commands should be executed
	baseDir string
	// Template for the deployment command
	cmdTemplate *template.Template
	// Package name to deploy
	packageName string
	// Dry run mode
	dryRun bool
	// Enclave name
	enclave string
}

const cmdTemplateStr = "just _kurtosis-run {{.PackageName}} {{.ArgFile}} {{.Enclave}}"

var defaultCmdTemplate *template.Template

func init() {
	defaultCmdTemplate = template.Must(template.New("kurtosis_deploy_cmd").Parse(cmdTemplateStr))
}

type KurtosisDeployerOptions func(*KurtosisDeployer)

func WithKurtosisBaseDir(baseDir string) KurtosisDeployerOptions {
	return func(d *KurtosisDeployer) {
		d.baseDir = baseDir
	}
}

func WithKurtosisCmdTemplate(cmdTemplate *template.Template) KurtosisDeployerOptions {
	return func(d *KurtosisDeployer) {
		d.cmdTemplate = cmdTemplate
	}
}

func WithKurtosisPackageName(packageName string) KurtosisDeployerOptions {
	return func(d *KurtosisDeployer) {
		d.packageName = packageName
	}
}

func WithKurtosisDryRun(dryRun bool) KurtosisDeployerOptions {
	return func(d *KurtosisDeployer) {
		d.dryRun = dryRun
	}
}

func WithKurtosisEnclave(enclave string) KurtosisDeployerOptions {
	return func(d *KurtosisDeployer) {
		d.enclave = enclave
	}
}

// NewKurtosisDeployer creates a new KurtosisDeployer instance
func NewKurtosisDeployer(opts ...KurtosisDeployerOptions) *KurtosisDeployer {
	d := &KurtosisDeployer{
		baseDir:     ".",
		cmdTemplate: defaultCmdTemplate,
		packageName: DefaultPackageName,
		dryRun:      false,
		enclave:     "devnet",
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

// templateData holds the data for the command template
type templateData struct {
	PackageName string
	ArgFile     string
	Enclave     string
}

// getInspectOutput runs kurtosis enclave inspect and writes its output to the provided writer
func (d *KurtosisDeployer) getInspectOutput(w io.Writer) error {
	cmd := exec.Command("kurtosis", "enclave", "inspect", d.enclave)
	cmd.Stdout = w
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to inspect kurtosis enclave: %w", err)
	}
	return nil
}

// findRPCEndpoint looks for a service matching the given predicate that has an RPC port
func findRPCEndpoints(services ServiceMap, matchService func(string) (string, bool)) EndpointMap {
	endpointMap := make(EndpointMap)
	for serviceName, ports := range services {
		if serviceName, ok := matchService(serviceName); ok {
			if port, ok := ports["rpc"]; ok {
				endpointMap[serviceName] = fmt.Sprintf("http://localhost:%d", port)
			}
			if port, ok := ports["http"]; ok {
				endpointMap[serviceName] = fmt.Sprintf("http://localhost:%d", port)
			}
		}
	}
	return endpointMap
}

func serviceTag(serviceName string) string {
	// Find index of first number
	i := strings.IndexFunc(serviceName, func(r rune) bool {
		return r >= '0' && r <= '9'
	})
	if i == -1 {
		return serviceName
	}
	return serviceName[:i-1]
}

// findL2RPCEndpoint looks for a service with the given suffix that has an RPC port
func findL2RPCEndpoints(services ServiceMap, suffix string) EndpointMap {
	return findRPCEndpoints(services, func(serviceName string) (string, bool) {
		if strings.HasSuffix(serviceName, suffix) {
			name := strings.TrimSuffix(serviceName, suffix)
			return strings.TrimPrefix(name, "op-"), true
		}
		return "", false
	})
}

// findL1RPCEndpoint looks the RPC port of the Ethereum service
func findL1RPCEndpoints(services ServiceMap) EndpointMap {
	return findRPCEndpoints(services, func(serviceName string) (string, bool) {
		match := !strings.HasPrefix(serviceName, "op-")
		if match {
			return serviceTag(serviceName), true
		}
		return "", false
	})
}

// prepareArgFile creates a temporary file with the input content and returns its path
func (d *KurtosisDeployer) prepareArgFile(input io.Reader) (string, error) {
	argFile, err := os.CreateTemp("", "kurtosis-args-*.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary arg file: %w", err)
	}
	defer argFile.Close()

	if _, err := io.Copy(argFile, input); err != nil {
		os.Remove(argFile.Name())
		return "", fmt.Errorf("failed to write arg file: %w", err)
	}

	return argFile.Name(), nil
}

// runKurtosisCommand executes the kurtosis command with the given arguments
func (d *KurtosisDeployer) runKurtosisCommand(argFile string) error {
	data := templateData{
		PackageName: d.packageName,
		ArgFile:     argFile,
		Enclave:     d.enclave,
	}

	var cmdBuf bytes.Buffer
	if err := d.cmdTemplate.Execute(&cmdBuf, data); err != nil {
		return fmt.Errorf("failed to execute command template: %w", err)
	}

	if d.dryRun {
		fmt.Println("Dry run mode enabled, kurtosis would run the following command:")
		fmt.Println(cmdBuf.String())
		return nil
	}

	cmd := exec.Command("sh", "-c", cmdBuf.String())
	cmd.Dir = d.baseDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("kurtosis deployment failed: %w", err)
	}

	return nil
}

// getEnvironmentInfo parses the input spec and inspect output to create KurtosisEnvironment
func (d *KurtosisDeployer) getEnvironmentInfo(spec *EnclaveSpec) (*KurtosisEnvironment, error) {
	var inspectBuf bytes.Buffer
	if err := d.getInspectOutput(&inspectBuf); err != nil {
		return nil, err
	}

	inspectResult, err := ParseInspectOutput(&inspectBuf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse inspect output: %w", err)
	}

	// Get contract addresses
	deployerState, err := ParseDeployerData(d.enclave)
	if err != nil {
		return nil, fmt.Errorf("failed to parse deployer state: %w", err)
	}

	env := &KurtosisEnvironment{
		L2: make([]*Chain, 0, len(spec.Chains)),
	}

	// Find L1 endpoint
	if rpcEndpoints := findL1RPCEndpoints(inspectResult.UserServices); len(rpcEndpoints) > 0 {
		env.L1 = &Chain{
			Name:      "Ethereum",
			Endpoints: rpcEndpoints,
		}
	}

	// Find L2 endpoints
	for _, chainSpec := range spec.Chains {
		rpcEndpoints := findL2RPCEndpoints(inspectResult.UserServices, chainSpec.Name)
		if len(rpcEndpoints) == 0 {
			continue
		}

		chain := &Chain{
			Name:      chainSpec.Name,
			ID:        chainSpec.NetworkID,
			Endpoints: rpcEndpoints,
		}

		// Add contract addresses if available
		if addresses, ok := deployerState[chainSpec.NetworkID]; ok {
			chain.Addresses = addresses
		}

		env.L2 = append(env.L2, chain)
	}

	return env, nil
}

// Deploy executes the Kurtosis deployment command with the provided input
func (d *KurtosisDeployer) Deploy(input io.Reader) (*KurtosisEnvironment, error) {
	// Parse the input spec first
	inputCopy := new(bytes.Buffer)
	tee := io.TeeReader(input, inputCopy)

	spec, err := ParseSpec(tee)
	if err != nil {
		return nil, fmt.Errorf("failed to parse input spec: %w", err)
	}

	// Prepare argument file
	argFile, err := d.prepareArgFile(inputCopy)
	if err != nil {
		return nil, err
	}
	defer os.Remove(argFile)

	// Run kurtosis command
	if err := d.runKurtosisCommand(argFile); err != nil {
		return nil, err
	}

	// If dry run, return empty environment
	if d.dryRun {
		return &KurtosisEnvironment{}, nil
	}

	// Get environment information
	return d.getEnvironmentInfo(spec)
}
