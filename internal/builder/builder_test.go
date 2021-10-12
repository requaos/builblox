package builder_test

import (
	"fmt"
	"github.com/requaos/builblox/internal/blox"
	"github.com/requaos/builblox/internal/builder"
	"github.com/requaos/builblox/internal/class"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBuild(t *testing.T) {
	testCases := [][]blox.Blox{{
		blox.New(class.Fire),
		blox.New(class.Spirit),
		blox.New(class.Water),
		blox.New(class.Wind),
		blox.New(class.Earth),
	},{
		blox.New(class.Spirit),
		blox.New(class.Wind),
		blox.New(class.Earth),
	},{
		blox.New(class.Water),
		blox.New(class.Wind),
		blox.New(class.Earth),
	},{
		blox.New(class.Fire),
		blox.New(class.Spirit),
		blox.New(class.Wind),
		blox.New(class.Earth),
	},{
		blox.New(class.Fire),
		blox.New(class.Spirit),
		blox.New(class.Water),
	},{
		blox.New(class.Fire),
		blox.New(class.Spirit),
		blox.New(class.Wind),
		blox.New(class.Spirit),
		blox.New(class.Wind),
		blox.New(class.Spirit),
		blox.New(class.Wind),
		blox.New(class.Spirit),
		blox.New(class.Wind),
		blox.New(class.Earth),
	}}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Build%d", i), func(t *testing.T) {
			b, err := testCase[0].Merge(testCase[1:]...)
			require.NoError(t, err)
			c := builder.Build(b)
			require.NotNil(t, c)
			t.Logf("\nType: %s\nLevel: %d", c.Name(), c.Level())
			b = c.Open()
			require.NotNil(t, b)
			list, err := b.Split()
			require.NoError(t, err)
			t.Log(list)
		})
	}
}
