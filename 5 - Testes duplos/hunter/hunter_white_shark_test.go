package hunter_test

import (
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"
)

func TestWhiteSharkHuntWithStubsAndMocks(t *testing.T) {
	mockSimulator := simulator.NewMockCatchSimulator(true)

	shark := hunter.CreateWhiteShark(mockSimulator)

	preyInstance := prey.NewStubPrey(50.0, &positioner.Position{X: 10, Y: 10, Z: 0})

	err := shark.Hunt(preyInstance)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if mockSimulator.CallCount() != 1 {
		t.Errorf("expected CanCatch to be called once, got %d", mockSimulator.CallCount())
	}
}

func TestWhiteSharkHuntWithFailedMock(t *testing.T) {
	mockSimulator := simulator.NewMockCatchSimulator(false)

	shark := hunter.CreateWhiteShark(mockSimulator)

	preyInstance := prey.NewStubPrey(200.0, &positioner.Position{X: 100, Y: 100, Z: 0})

	err := shark.Hunt(preyInstance)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}

	if mockSimulator.CallCount() != 1 {
		t.Errorf("expected CanCatch to be called once, got %d", mockSimulator.CallCount())
	}
}
