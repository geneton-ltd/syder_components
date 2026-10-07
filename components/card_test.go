package components_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestCard_Options(t *testing.T) {
	// create card component with options
	comp := co.Card(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func hasCardClasses(t *testing.T, class string) {
	require.Contains(t, class, "w-full")
	require.Contains(t, class, "rounded-md")
	require.Contains(t, class, "shadow-xs")
	require.Contains(t, class, "p-2")
	require.Contains(t, class, "border")
	require.Contains(t, class, "border-1")
	require.Contains(t, class, "border-gray-600")
}

func TestCard_DefaultStyle(t *testing.T) {
	comp := co.Card()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasCardClasses(t, class)
}

func TestCard_Output(t *testing.T) {
	id := "id"
	comp := co.Card(sc.ID(id))
	tu.AssertElementWithIDExists(t, comp, "div", id)
}

func TestSubCard_Merging(t *testing.T) {
	comp := co.SubCard(sc.Class("custom-class"))
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)

	// has the default Card style
	hasCardClasses(t, class)

	// has its own style
	require.Contains(t, class, "border-dashed")
	require.Contains(t, class, "hover:bg-white")

	// has custom class
	require.Contains(t, class, "custom-class")
}
