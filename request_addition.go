package main


type RequestAdditon struct {
	validators []RequestValidator
	registeredList []RegisterRequest
}

func NewRequestAdditon(liftCapacity, maxStops int) *RequestAdditon {
	return &RequestAdditon{
		validators: []RequestValidator{
			NewCapacityRequestValidator(liftCapacity),
			NewMaxStopRequestValidation(maxStops),
		},
		registeredList: make([]RegisterRequest, 0),
	}
}

func (ra *RequestAdditon) AddRequest(request RegisterRequest) bool {
	for _,validator := range ra.validators {
		if !validator.ValidateRequest(ra.registeredList, request){
			return false
		}
	}
	ra.registeredList=append(ra.registeredList,request)
    return true
}