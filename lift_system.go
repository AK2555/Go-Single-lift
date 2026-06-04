package main

type LiftSystem struct {
	directionScheduler *DirectionScheduler
	currentState LiftState
}

func NewLiftSystem(floorsCount, liftCapacity, maxStops int) *LiftSystem {
	ls := &LiftSystem{
		directionScheduler: NewDirectionScheduler(liftCapacity,maxStops),
	}
	ls.currentState=NewIdleState(ls)
	return ls
}

func (ls *LiftSystem) RequestPickup(source, destination int) bool {
	return ls.currentState.HandleRequest(source,destination)
}

func (ls *LiftSystem) GetDirectionScheduler() *DirectionScheduler {
	return ls.directionScheduler
}