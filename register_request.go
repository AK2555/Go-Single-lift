package main

type RegisterRequest struct {
	source int
	destination int
	direction Direction
}

func NewRegisterRequest(source , destination int,direction Direction) RegisterRequest {
	return RegisterRequest{
		source: source,
		destination: destination,
		direction: direction,
	}
}

func (r *RegisterRequest) SetSource(source int){
	r.source=source
}

func (r RegisterRequest) GetSource() int {
	return r.source
}

func (r *RegisterRequest) SetDestination(destination int){
	r.destination=destination
}

func (r RegisterRequest) GetDestination() int {
	return r.destination
}

func (r *RegisterRequest) SetDirection(direction Direction){
    r.direction=direction
}

func (r RegisterRequest) GetDirection() Direction {
	return r.direction
}