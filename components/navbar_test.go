package components_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestNavbar_Options(t *testing.T) {
	// create navbar component with options
	comp := co.Navbar(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func TestNavbar_DefaultStyle(t *testing.T) {
	comp := co.Navbar()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	require.Contains(t, class, "h-[50px]")
	require.Contains(t, class, "bg-gray-800")
	require.Contains(t, class, "flex")
	require.Contains(t, class, "justify-between")
	require.Contains(t, class, "text-white")
}

func TestNavbar_Output(t *testing.T) {
	id := "id"
	comp := co.Navbar(sc.ID(id))
	tu.AssertElementWithIDExists(t, comp, "div", id)
}

func TestNavLink_Options(t *testing.T) {
	// create navlink component with options
	comp := co.NavLink(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func TestNavLink_DefaultStyle(t *testing.T) {
	comp := co.NavLink()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	require.Contains(t, class, "hover:bg-gray-600")
	require.Contains(t, class, "flex")
	require.Contains(t, class, "items-center")
	require.Contains(t, class, "px-2")
}

func TestNavLink_Output(t *testing.T) {
	id := "id"
	comp := co.NavLink(sc.ID(id))
	tu.AssertElementWithIDExists(t, comp, "a", id)
}
