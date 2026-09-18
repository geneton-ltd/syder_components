package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestPopover_Options(t *testing.T) {
	// create popover component with options
	comp := sc.Popover(sc.WithID("id"), sc.WithData("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestPopover_DefaultStyle(t *testing.T) {
	comp := sc.Popover()
	class := comp.GetClass()

	require.Contains(t, class, "shadow-xl")
	require.Contains(t, class, "bg-white")
	require.Contains(t, class, "rounded-md")
	require.Contains(t, class, "absolute")
	require.Contains(t, class, "border")
	require.Contains(t, class, "p-4")

	attrs := comp.GetAttrs()
	popoverValue, ok := attrs["popover"]
	require.True(t, ok)
	require.Equal(t, "auto", popoverValue)
}

func TestPopover_Output(t *testing.T) {
	id := "id"
	comp := sc.Popover(sc.WithID(id))
	assertElementWithIDExists(t, comp, "div", id)
}
