package main

type LiftState interface {
	HandleRequest(source,destination int) bool
}

type IdleState struct {
	liftSystem *LiftSystem
}

func NewIdleState(liftSystem *LiftSystem) *IdleState {
	return &IdleState{liftSystem: liftSystem}
}

func (s *IdleState) HandleRequest(source, destination int) bool {
	direction := UP
	if source > destination {
		direction = DOWN
	}
	req := NewRegisterRequest(source,destination,direction)
	return s.liftSystem.GetDirectionScheduler().GetRequestAdditon(direction).AddRequest(req)
}