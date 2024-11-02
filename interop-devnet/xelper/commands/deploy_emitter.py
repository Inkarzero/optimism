import subprocess
import argparse
import json
from util import RPC_URLs, MONOREPO_ROOT

description = "List the contents of a directory and save to context."

def execute(context, args):
  parser = argparse.ArgumentParser()
  parser.add_argument('chain', type=str, help="Chain to deploy the emitter contract to")
  args = parser.parse_args(args)

  rpc_url = RPC_URLs[args.chain]
  root = MONOREPO_ROOT

  cmd = rf"""
  #!/bin/bash
  cd {root}
  OP_INTEROP_MNEMONIC="test test test test test test test test test test test junk"
  OP_INTEROP_DEVKEY_CHAINID=0
  OP_INTEROP_DEVKEY_DOMAIN=user
  OP_INTEROP_DEVKEY_NAME=0
  ETH_RPC_URL={rpc_url}
  RAW_PRIVATE_KEY=0x$(go run ./op-node/cmd interop devkey secret)
  cd op-e2e/e2eutils/interop/contracts/

  # TODO: make forge build its own command, this takes too much time
  #forge build 2&>1 /dev/null

  forge create --private-key=$RAW_PRIVATE_KEY "src/emit.sol:EmitEvent" --json
  """

  result = subprocess.run(cmd, capture_output=True, executable='/bin/bash', shell=True, text=True)
  resultJSON = json.loads(result.stdout)

  context.set(f"{args.chain}.CreateEmitterOutput", resultJSON)
  context.set(f"{args.chain}.EmitterContractAddress", resultJSON["deployedTo"])

  print(f"Emitter contract created: {resultJSON}")
  return(result.stdout, result.stderr)