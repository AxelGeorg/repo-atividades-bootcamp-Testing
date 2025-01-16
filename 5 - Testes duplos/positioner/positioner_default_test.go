package positioner_test

import (
	"testdoubles/positioner"
	"testing"
)

func TestStubPositioner(t *testing.T) {
	stub := &positioner.StubPositioner{
		DistanceToReturn: 10.0,
	}

	tests := []struct {
		name     string
		from     *positioner.Position
		to       *positioner.Position
		expected float64
	}{
		{
			name:     "Test case with any positions",
			from:     &positioner.Position{X: 0, Y: 0, Z: 0},
			to:       &positioner.Position{X: 1, Y: 1, Z: 1},
			expected: 10.0, // A unidade de distância predefinida que você definiu no stub
		},
		{
			name:     "Same positions",
			from:     &positioner.Position{X: 2, Y: 2, Z: 2},
			to:       &positioner.Position{X: 2, Y: 2, Z: 2},
			expected: 10.0, // Mesmo aqui, o retorno é constante
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stub.GetLinearDistance(tt.from, tt.to)
			if result != tt.expected {
				t.Errorf("expected distance %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestGetLinearDistance(t *testing.T) {
	tests := []struct {
		number   int
		name     string
		from     *positioner.Position
		to       *positioner.Position
		expected float64
	}{
		{
			number:   1,
			name:     "All coordinates are negative",
			from:     &positioner.Position{X: -3, Y: -4, Z: 0},
			to:       &positioner.Position{X: -6, Y: -8, Z: 0},
			expected: 5.0, // Distância esperada é a mesma da coordenada positiva (5.0)
		},
		{
			number:   2,
			name:     "All coordinates are positive",
			from:     &positioner.Position{X: 3, Y: 4, Z: 0},
			to:       &positioner.Position{X: 6, Y: 8, Z: 0},
			expected: 5.0, // Distância esperada
		},
		{
			number:   3,
			name:     "Distance without decimals",
			from:     &positioner.Position{X: 1, Y: 1, Z: 1},
			to:       &positioner.Position{X: 4, Y: 1, Z: 1},
			expected: 3.0, // Distância esperada, pois se move apenas ao longo do eixo X
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stub *positioner.StubPositioner

			if tt.number == 1 || tt.number == 2 {
				stub = &positioner.StubPositioner{
					DistanceToReturn: 5.0,
				}
			} else if tt.number == 3 {
				stub = &positioner.StubPositioner{
					DistanceToReturn: 3.0,
				}
			}

			result := stub.GetLinearDistance(tt.from, tt.to)
			if result != tt.expected {
				t.Errorf("expected distance %v, got %v", tt.expected, result)
			}
		})
	}
}
