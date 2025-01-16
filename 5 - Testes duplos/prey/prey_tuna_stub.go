package prey

import "testdoubles/positioner"

type StubPrey struct {
	speed    float64
	position *positioner.Position
}

func NewStubPrey(speed float64, position *positioner.Position) *StubPrey {
	return &StubPrey{
		speed:    speed,
		position: position,
	}
}

func (s *StubPrey) GetSpeed() float64 {
	return s.speed
}

func (s *StubPrey) GetPosition() *positioner.Position {
	return s.position
}
