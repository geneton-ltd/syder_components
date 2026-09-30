package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestAny_Options(t *testing.T) {
	// create any component with options
	comp := sc.Any("div", sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
	require.Equal(t, "div", comp.GetTag())
}

func TestAny_Output(t *testing.T) {
	id := "id"

	tags := []string{"button", "a", "div", "input", "p"}
	for _, tag := range tags {
		comp := sc.Any(tag, sc.ID(id))
		assertElementWithIDExists(t, comp, tag, id)
		require.Equal(t, tag, comp.GetTag())
	}
}
