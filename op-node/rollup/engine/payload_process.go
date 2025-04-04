package engine

import (
	"context"
	"fmt"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type PayloadProcessEvent struct {
	// if payload should be promoted to safe (must also be pending safe, see DerivedFrom)
	IsLastInSpan bool
	// payload is promoted to pending-safe if non-zero
	DerivedFrom eth.L1BlockRef

	Envelope *eth.ExecutionPayloadEnvelope
	Ref      eth.L2BlockRef
}

func (ev PayloadProcessEvent) String() string {
	return "payload-process"
}

func (eq *EngDeriver) onPayloadProcess(ev PayloadProcessEvent) {
	eq.log.Debug("/op-node/rollup/engine/payload_process.go | onPayloadProcess | started", "ev", ev)
	ctx, cancel := context.WithTimeout(eq.ctx, payloadProcessTimeout)
	defer cancel()

	status, err := eq.ec.engine.NewPayload(ctx,
		ev.Envelope.ExecutionPayload, ev.Envelope.ParentBeaconBlockRoot)
	if err != nil {
		eq.log.Debug("/op-node/rollup/engine/payload_process.go | onPayloadProcess | emitting EngineTemporaryErrorEvent", "err", err)
		eq.emitter.Emit(rollup.EngineTemporaryErrorEvent{
			Err: fmt.Errorf("failed to insert execution payload: %w", err)})
		return
	}
	switch status.Status {
	case eth.ExecutionInvalid, eth.ExecutionInvalidBlockHash:
		eq.log.Debug("/op-node/rollup/engine/payload_process.go | onPayloadProcess | invalid payload", "status", status)
		eq.emitter.Emit(PayloadInvalidEvent{
			Envelope: ev.Envelope,
			Err:      eth.NewPayloadErr(ev.Envelope.ExecutionPayload, status)})
		return
	case eth.ExecutionValid:
		eq.log.Debug("/op-node/rollup/engine/payload_process.go | onPayloadProcess | valid payload", "status", status)
		eq.emitter.Emit(PayloadSuccessEvent(ev))
		return
	default:
		eq.log.Debug("/op-node/rollup/engine/payload_process.go | onPayloadProcess | temporary error", "status", status)
		eq.emitter.Emit(rollup.EngineTemporaryErrorEvent{
			Err: eth.NewPayloadErr(ev.Envelope.ExecutionPayload, status)})
		return
	}
}
