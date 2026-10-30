package sydercomponents

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	options := []string{"a", "b", "c", "d"}
	comp := Select(Map(options, func(opt string) Node { return Option(Value(opt), Text(opt)) }))

	// has |options| children and each has mapped structure
	require.Equal(t, len(options), len(comp.Children()))
	for idx, child := range comp.Children() {
		require.Equal(t, "option", child.Name())
		value, ok := child.Attr("value")
		require.True(t, ok)
		require.Equal(t, options[idx], value)
		require.Len(t, child.Children(), 1)
		require.Equal(t, KindText, child.Children()[0].Kind())
		require.Equal(t, options[idx], child.Children()[0].Value())
	}
}
