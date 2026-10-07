package components_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestDialog_Options(t *testing.T) {
	// create dialog component with options
	comp := sc.Dialog(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func hasDialogClasses(t *testing.T, class string) {
	require.Contains(t, class, "shadow-xl")
	require.Contains(t, class, "bg-white")
	require.Contains(t, class, "backdrop:bg-black/40")
	require.Contains(t, class, "rounded-md")
}

func TestDialog_DefaultStyle(t *testing.T) {
	comp := sc.Dialog()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasDialogClasses(t, class)
}

func TestDialog_Output(t *testing.T) {
	id := "id"
	comp := sc.Dialog(sc.ID(id))
	tu.AssertElementWithIDExists(t, comp, "dialog", id)
}

func TestDialogRelative_Merging(t *testing.T) {
	comp := co.DialogRelative(sc.Class("custom-class"))
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)

	// has the default Dialog style
	hasDialogClasses(t, class)

	// has its own style
	require.Contains(t, class, "absolute")
	require.Contains(t, class, "m-0")

	// has custom class
	require.Contains(t, class, "custom-class")
}

func TestDialogFixedCentered_Merging(t *testing.T) {
	comp := co.DialogFixedCentered(sc.Class("custom-class"))
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)

	// has the default Dialog style
	hasDialogClasses(t, class)

	// has its own style
	require.Contains(t, class, "fixed")
	require.Contains(t, class, "inset-0")
	require.Contains(t, class, "m-auto")

	// has custom class
	require.Contains(t, class, "custom-class")
}
