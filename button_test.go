package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestButton_Options(t *testing.T) {
	// create button component with options
	comp := sc.Button(sc.WithID("id"), sc.WithData("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasButtonClasses(t *testing.T, class string) {
	require.Contains(t, class, "rounded")
	require.Contains(t, class, "rounded-sm")
	require.Contains(t, class, "shadow-xs")
	require.Contains(t, class, "p-1")
	require.Contains(t, class, "hover:cursor-pointer")
	require.Contains(t, class, "disabled:opacity-50")
	require.Contains(t, class, "disabled:cursor-not-allowed")
}

func TestButton_DefaultStyle(t *testing.T) {
	comp := sc.Card()
	hasCardClasses(t, comp.GetClass())
}

func TestButton_Output(t *testing.T) {
	id := "id"
	comp := sc.Button(sc.WithID(id))
	assertElementWithIDExists(t, comp, "button", id)
}

func hasIconButtonClasses(t *testing.T, class string) {
	require.Contains(t, class, "whitespace-nowrap")
	require.Contains(t, class, "flex")
	require.Contains(t, class, "gap-1")
	require.Contains(t, class, "items-center")
}

func TestIconButton_Merging(t *testing.T) {
	comp := sc.IconButton(sc.Any("div"), sc.WithClassAppend("custom-class"))
	class := comp.GetClass()

	// has the default Button style
	hasButtonClasses(t, class)

	// has its own style
	hasIconButtonClasses(t, class)

	// has custom class
	require.Contains(t, class, "custom-class")

	// has one child - <span>
	children := comp.GetChildren()
	require.Len(t, children, 1)

	iconComp := children[0]
	require.Equal(t, "h-4 w-4 shrink-0", iconComp.GetClass())
}

func TestIconOnlyButton_Merging(t *testing.T) {
	comp := sc.IconOnlyButton(sc.Any("div"), sc.WithClassAppend("custom-class"))
	class := comp.GetClass()

	// has the default Button style
	hasButtonClasses(t, class)

	// has the default IconButton style
	hasIconButtonClasses(t, class)

	// has its own style
	require.Contains(t, class, "w-fit")
	require.Contains(t, class, "self-center")

	// has custom class
	require.Contains(t, class, "custom-class")

	// is type="button"
	compType, hasType := comp.GetType()
	require.True(t, hasType)
	require.Equal(t, "button", compType)
}
