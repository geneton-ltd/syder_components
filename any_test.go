package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestAny_Options(t *testing.T) {
	// create any component with options
	comp := sc.Any("div", sc.WithID("id"), sc.WithData("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestAny_Output(t *testing.T) {
	id := "id"

	tags := []string{"button", "a", "div", "input", "p"}
	for _, tag := range tags {
		comp := sc.Any(tag, sc.WithID(id))
		assertElementWithIDExists(t, comp, tag, id)
	}
}
