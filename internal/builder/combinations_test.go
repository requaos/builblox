package builder_test

import (
	"math/rand"
	"testing"

	"github.com/requaos/builblox/internal/builder"
	"github.com/stretchr/testify/require"
)

func TestCombinations(t *testing.T) {
	t.Parallel()

	t.Run("fixed_5", func(t *testing.T) {
		rand.Seed(3)
		have := builder.TokenIDs(5)
		//have := []string{"test", "testing", "tester", "bling", "bang"}
		want := 31

		n, combos := builder.Combinations(have)
		t.Log(n)
		for _, combo := range combos {
			t.Log(combo)
		}

		require.True(t, n == want)
		require.Len(t, combos, want)
	})

	t.Run("dynamic_20", func(t *testing.T) {
		rand.Seed(33)
		n:=20
		have := builder.TokenIDs(n)
		want := builder.TotalCombinationsFor(n)

		n, combos := builder.Combinations(have)
		t.Log(n)

		require.True(t, n == want)
		require.Len(t, combos, want)
	})
}

func TestTokenIDs(t *testing.T) {
	t.Parallel()

	testCount := 5
	out := builder.TokenIDs(testCount)
	for _, token := range out {
		t.Log(token)
	}
	require.Len(t, out, testCount)
}
