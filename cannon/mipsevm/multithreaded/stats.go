package multithreaded

import (
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm"
)

// Define stats interface
type StatsTracker interface {
	trackLL(step uint64)
	trackSCSuccess(step uint64)
	trackSCFailure(step uint64)
	trackReservationInvalidation()
	trackForcedPreemption()
	trackWakeupTraversalStart()
	trackWakeup()
	trackWakeupFail()
	populateDebugInfo(debugInfo *mipsevm.DebugInfo)
}

// Noop implementation for when debug is disabled
type noopStatsTracker struct{}

func NoopStatsTracker() StatsTracker {
	return &noopStatsTracker{}
}

func (s *noopStatsTracker) trackLL(step uint64)                            {}
func (s *noopStatsTracker) trackSCSuccess(step uint64)                     {}
func (s *noopStatsTracker) trackSCFailure(step uint64)                     {}
func (s *noopStatsTracker) trackReservationInvalidation()                  {}
func (s *noopStatsTracker) trackForcedPreemption()                         {}
func (s *noopStatsTracker) trackWakeupTraversalStart()                     {}
func (s *noopStatsTracker) trackWakeup()                                   {}
func (s *noopStatsTracker) trackWakeupFail()                               {}
func (s *noopStatsTracker) populateDebugInfo(debugInfo *mipsevm.DebugInfo) {}

var _ StatsTracker = (*noopStatsTracker)(nil)

// Actual implementation
type statsTrackerImpl struct {
	// State
	lastLLOpStep      uint64
	isWakeupTraversal bool
	// Stats
	rmwSuccessCount int
	rmwFailCount    int
	// Note: Once a new LL operation is executed, we reset lastLLOpStep, losing track of previous RMW operations.
	// So, maxStepsBetweenLLAndSC is not complete and may miss longer ranges for failed rmw sequences.
	maxStepsBetweenLLAndSC uint64
	// Tracks RMW reservation invalidation due to reserved memory being accessed outside of the RMW sequence
	reservationInvalidationCount int
	forcedPreemptionCount        int
	failedWakeupCount            int
}

func (s *statsTrackerImpl) populateDebugInfo(debugInfo *mipsevm.DebugInfo) {
	debugInfo.RmwSuccessCount = s.rmwSuccessCount
	debugInfo.RmwFailCount = s.rmwFailCount
	debugInfo.MaxStepsBetweenLLAndSC = hexutil.Uint64(s.maxStepsBetweenLLAndSC)
	debugInfo.ReservationInvalidationCount = s.reservationInvalidationCount
	debugInfo.ForcedPreemptionCount = s.forcedPreemptionCount
	debugInfo.FailedWakeupCount = s.failedWakeupCount
}

func (s *statsTrackerImpl) trackLL(step uint64) {
	s.lastLLOpStep = step
}

func (s *statsTrackerImpl) trackSCSuccess(step uint64) {
	s.rmwSuccessCount += 1
	diff := step - s.lastLLOpStep
	if diff > s.maxStepsBetweenLLAndSC {
		s.maxStepsBetweenLLAndSC = diff
	}
	// Reset ll op state
	s.lastLLOpStep = 0
}

func (s *statsTrackerImpl) trackSCFailure(step uint64) {
	s.rmwFailCount += 1

	diff := step - s.lastLLOpStep
	if s.lastLLOpStep > 0 && diff > s.maxStepsBetweenLLAndSC {
		s.maxStepsBetweenLLAndSC = diff
	}
}

func (s *statsTrackerImpl) trackReservationInvalidation() {
	s.reservationInvalidationCount += 1
}

func (s *statsTrackerImpl) trackForcedPreemption() {
	s.forcedPreemptionCount += 1
}

func (s *statsTrackerImpl) trackWakeupTraversalStart() {
	s.isWakeupTraversal = true
}

func (s *statsTrackerImpl) trackWakeup() {
	s.isWakeupTraversal = false
}

func (s *statsTrackerImpl) trackWakeupFail() {
	if s.isWakeupTraversal {
		s.failedWakeupCount += 1
	}
	s.isWakeupTraversal = false
}

func NewStatsTracker() StatsTracker {
	return &statsTrackerImpl{}
}

var _ StatsTracker = (*statsTrackerImpl)(nil)
