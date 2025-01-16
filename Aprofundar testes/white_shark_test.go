package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 9000)
		err := shark.Hunt(&hunt.Tuna{Speed: 800})
		require.NoError(t, err)
		require.False(t, shark.Hungry)
		require.True(t, shark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		shark := hunt.NewWhiteShark(false, false, 9000)
		err := shark.Hunt(&hunt.Tuna{Speed: 800})
		require.ErrorIs(t, err, hunt.ErrSharkIsNotHungry)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, true, 9000)
		err := shark.Hunt(&hunt.Tuna{Speed: 800})
		require.ErrorIs(t, err, hunt.ErrSharkIsTired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 9000)
		err := shark.Hunt(&hunt.Tuna{Speed: 10000})
		require.ErrorIs(t, err, hunt.ErrSharkIsSlower)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, true, 9000)
		err := shark.Hunt(nil)
		require.ErrorIs(t, err, hunt.ErrNoTuna)
	})
}
