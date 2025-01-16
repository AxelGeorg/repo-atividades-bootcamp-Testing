package simulator

type MockCatchSimulator struct {
	canCatchResult bool
	callCount      int
}

func NewMockCatchSimulator(canCatchResult bool) *MockCatchSimulator {
	return &MockCatchSimulator{
		canCatchResult: canCatchResult,
	}
}

func (m *MockCatchSimulator) CanCatch(hunter, prey *Subject) bool {
	m.callCount++
	return m.canCatchResult
}

func (m *MockCatchSimulator) CallCount() int {
	return m.callCount
}

func (m *MockCatchSimulator) Reset() {
	m.callCount = 0
}
