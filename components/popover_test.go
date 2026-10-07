package components_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestPopover_Options(t *testing.T) {
	// create popover component with options
	comp := co.Popover(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func TestPopover_DefaultStyle(t *testing.T) {
	comp := co.Popover()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)

	require.Contains(t, class, "shadow-xl")
	require.Contains(t, class, "bg-white")
	require.Contains(t, class, "rounded-md")
	require.Contains(t, class, "absolute")
	require.Contains(t, class, "border")
	require.Contains(t, class, "p-4")

	popoverValue, ok := comp.Attr("popover")
	require.True(t, ok)
	require.Equal(t, "auto", popoverValue)
}

func TestPopover_Output(t *testing.T) {
	id := "id"
	comp := co.Popover(sc.ID(id))
	tu.AssertElementWithIDExists(t, comp, "div", id)
}
