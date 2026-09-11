package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestDialog_Options(t *testing.T) {
	// create dialog component with options
	comp := sc.Dialog(sc.WithID("id"), sc.WithData("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasDialogClasses(t *testing.T, class string) {
	require.Contains(t, class, "shadow-xl")
	require.Contains(t, class, "bg-white")
	require.Contains(t, class, "backdrop:bg-black/40")
	require.Contains(t, class, "rounded-md")
}

func TestDialog_DefaultStyle(t *testing.T) {
	comp := sc.Dialog()

	hasDialogClasses(t, comp.GetClass())
}

func TestDialog_Output(t *testing.T) {
	id := "id"
	comp := sc.Dialog(sc.WithID(id))
	assertElementWithIDExists(t, comp, "dialog", id)
}

func TestDialogRelative_Merging(t *testing.T) {
	comp := sc.DialogRelative(sc.WithClassAppend("custom-class"))
	class := comp.GetClass()

	// has the default Dialog style
	hasDialogClasses(t, class)

	// has its own style
	require.Contains(t, class, "absolute")
	require.Contains(t, class, "m-0")

	// has custom class
	require.Contains(t, class, "custom-class")
}

func TestDialogFixedCentered_Merging(t *testing.T) {
	comp := sc.DialogFixedCentered(sc.WithClassAppend("custom-class"))
	class := comp.GetClass()

	// has the default Dialog style
	hasDialogClasses(t, class)

	// has its own style
	require.Contains(t, class, "fixed")
	require.Contains(t, class, "inset-0")
	require.Contains(t, class, "m-auto")

	// has custom class
	require.Contains(t, class, "custom-class")
}
