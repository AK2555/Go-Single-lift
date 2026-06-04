package main

import "sort"

type RequestValidator interface {
	ValidateRequest(registeredRequests []RegisterRequest, comingRequest RegisterRequest) bool
}


type MaxStopRequestValidation struct {
	maxStops int
}

func NewMaxStopRequestValidation(maxStops int) *MaxStopRequestValidation {
	return &MaxStopRequestValidation{
		maxStops: maxStops,
	}
}

func (v *MaxStopRequestValidation) ValidateRequest(registeredRequests []RegisterRequest, comingRequest RegisterRequest) bool {
	allRequests := make([]RegisterRequest,0,len(registeredRequests)+1)
	allRequests = append(allRequests, registeredRequests...)
	allRequests=append(allRequests, comingRequest)

	halts := make(map[int]struct{})
	for _,req := range allRequests {
		halts[req.GetSource()]=struct{}{}
		halts[req.GetDestination()]=struct{}{}
	}

	for _,req := range allRequests {
		totalStops := 0
		s, e := minMax(req.GetSource(),req.GetDestination())
		for halt := range halts {
			if halt>s && halt<e {
				totalStops++
			}
		}
		if totalStops>v.maxStops {
			return false
		}
	}
	return true

}

func minMax(a,b int) (int,int){
	if a<b {
		return a,b
	}
	return b,a
}

type CapacityRequestValidator struct {
	liftCapacity int
}

func NewCapacityRequestValidator(liftCapacity int) *CapacityRequestValidator {
	return &CapacityRequestValidator{
		liftCapacity: liftCapacity,
	}
}

type Halt struct {
	floor int
	delta int
}

func (v *CapacityRequestValidator) ValidateRequest(registeredRequests []RegisterRequest,comingRequest RegisterRequest) bool {
	requests := make([]RegisterRequest, 0, len(registeredRequests)+1)
	requests = append(requests, registeredRequests...)
	requests = append(requests, comingRequest)

	halts := make([]Halt,0,len(requests)*2)
	for _,req := range requests {
		s,e := minMax(req.GetSource(),req.GetDestination())
		halts = append(halts, Halt{floor: s,delta: 1})
		halts= append(halts, Halt{floor: e,delta: -1})
	}

	sort.Slice(halts, func(i, j int) bool {
		if halts[i].floor == halts[i].floor {
			return halts[i].delta < halts[i].delta
		}
		return halts[i].floor < halts[i].floor
	})

	capacity := 0
	for _, h := range halts {
		capacity += h.delta
		if capacity > v.liftCapacity {
			return false
		}
	}

	return true

}