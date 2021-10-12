package class_test

import (
	"github.com/requaos/builblox/internal/class"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestType(t *testing.T) {
	t.Parallel()

	t.Run("NewSuccess", func(t *testing.T) {
		v, err := class.New("fire")
		require.NoError(t, err)
		require.NotNil(t, v)
		require.Equal(t, "fire", v.String())
	})

	t.Run("NewFailed", func(t *testing.T) {
		v, err := class.New("test")
		require.EqualError(t, err, class.ErrInvalidType.Error())
		require.Nil(t, v)
	})
}
