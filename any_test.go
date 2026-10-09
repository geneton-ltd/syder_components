package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestAny_Options(t *testing.T) {
	// create any component with options
	comp := sc.Any("div", sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.Equal(t, "id", v)
	require.True(t, has)

	v, has = comp.Attr("data-name")
	require.Equal(t, "john", v)
	require.True(t, has)
	require.Equal(t, sc.KindElement, comp.Kind())
}

func TestAny_Output(t *testing.T) {
	id := "id"

	tags := []string{"button", "a", "div", "input", "p"}
	for _, tag := range tags {
		comp := sc.Any(tag, sc.ID(id))
		tu.AssertElementWithIDExists(t, comp, tag, id)
		require.Equal(t, sc.KindElement, comp.Kind())
	}
}
