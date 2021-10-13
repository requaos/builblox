package blox_test

import (
	"testing"

	"github.com/requaos/builblox/internal/blox"
	"github.com/requaos/builblox/internal/class"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Parallel()

	b := blox.New(class.Fire)
	require.NotNil(t, b)
	require.Equal(t, class.Fire.String(), b.String())
	found := false
	for name, count := range b.Contents() {
		if name == class.Fire && count == 1 {
			found = true
		}
	}
	require.True(t, found)
}

var testCases = []class.Type{
	class.Fire, class.Fire,
	class.Wind,
	class.Water, class.Water}

func TestSplit(t *testing.T) {
	t.Parallel()

	tc := make([]blox.Blox, 0, len(testCases))
	for _, name := range testCases {
		tc = append(tc, blox.New(name))
	}

	b, err := tc[0].Merge(tc[1:]...)
	require.NoError(t, err)

	all, err := b.Split()
	require.NoError(t, err)
	require.Len(t, all, len(testCases))
}

func TestMerge(t *testing.T) {
	t.Parallel()

	tc := make([]blox.Blox, 0, len(testCases))
	for _, name := range testCases {
		tc = append(tc, blox.New(name))
	}

	b, err := tc[0].Merge(tc[1:]...)
	require.NoError(t, err)

	counts := 0
	contents := b.Contents()
	for _, count := range contents {
		counts += count
	}
	require.Equal(t, len(tc), counts)
}

func TestDestroy(t *testing.T) {
	t.Parallel()

	testCase := class.Fire
	b := blox.New(testCase)
	require.Equal(t, 1, b.Contents()[testCase])

	b.Destroy()

	_, ok := b.Contents()[testCase]
	require.False(t, ok)
}
