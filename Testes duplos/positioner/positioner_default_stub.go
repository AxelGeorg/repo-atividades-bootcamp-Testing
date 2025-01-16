package positioner

type StubPositioner struct {
	DistanceToReturn float64
}

func (s *StubPositioner) GetLinearDistance(from, to *Position) float64 {
	return s.DistanceToReturn
}
