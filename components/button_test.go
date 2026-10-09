package components_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestButton_Options(t *testing.T) {
	// create button component with options
	comp := co.Button(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
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
	comp := co.Button()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasButtonClasses(t, class)
}

func TestButton_Output(t *testing.T) {
	id := "id"
	comp := co.Button(sc.ID(id))
	tu.AssertElementWithIDExists(t, comp, "button", id)
}

func hasIconButtonClasses(t *testing.T, class string) {
	require.Contains(t, class, "whitespace-nowrap")
	require.Contains(t, class, "flex")
	require.Contains(t, class, "gap-1")
	require.Contains(t, class, "items-center")
}

func TestIconButton_Merging(t *testing.T) {
	comp := co.IconButton(sc.Div(), sc.Class("custom-class"))
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)

	// has the default Button style
	hasButtonClasses(t, class)

	// has its own style
	hasIconButtonClasses(t, class)

	// has custom class
	require.Contains(t, class, "custom-class")

	// has one child - <span>
	children := comp.Children()
	require.Len(t, children, 1)

	iconComp := children[0]
	class, hasClass = iconComp.Attr("class")
	require.True(t, hasClass)
	require.Equal(t, "h-4 w-4 shrink-0", class)
}

func TestIconOnlyButton_Merging(t *testing.T) {
	comp := co.IconOnlyButton(sc.Div(), sc.Class("custom-class"))
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)

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
	compType, hasType := comp.Attr("type")
	require.True(t, hasType)
	require.Equal(t, "button", compType)
}
