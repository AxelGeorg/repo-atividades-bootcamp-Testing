package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"
)

func TestStubPrey(t *testing.T) {
	tests := []struct {
		name     string
		speed    float64
		position *positioner.Position
	}{
		{
			name:  "Fast prey",
			speed: 200.0,
			position: &positioner.Position{
				X: 100,
				Y: 100,
				Z: 0,
			},
		},
		{
			name:  "Slow prey",
			speed: 50.0,
			position: &positioner.Position{
				X: 200,
				Y: 200,
				Z: 0,
			},
		},
		{
			name:  "Stationary prey",
			speed: 0.0,
			position: &positioner.Position{
				X: 300,
				Y: 300,
				Z: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stubPrey := prey.NewStubPrey(tt.speed, tt.position)

			if got := stubPrey.GetSpeed(); got != tt.speed {
				t.Errorf("GetSpeed() = %v; want %v", got, tt.speed)
			}

			if got := stubPrey.GetPosition(); got.X != tt.position.X || got.Y != tt.position.Y || got.Z != tt.position.Z {
				t.Errorf("GetPosition() = %v; want %v", got, tt.position)
			}
		})
	}
}
