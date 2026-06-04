package main

type DirectionScheduler struct {
	m map[Direction]*RequestAdditon
}

func NewDirectionScheduler(liftCapacity, maxStops int) *DirectionScheduler {
	ds := &DirectionScheduler{
		m: make(map[Direction]*RequestAdditon),
	}
	ds.m[UP] = NewRequestAdditon(liftCapacity, maxStops)
	ds.m[DOWN] = NewRequestAdditon(liftCapacity, maxStops)
	return ds
}

func (ds *DirectionScheduler) GetRequestAdditon(direction Direction) *RequestAdditon {
	return ds.m[direction]
}