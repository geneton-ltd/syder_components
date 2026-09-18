package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestNavbar_Options(t *testing.T) {
	// create navbar component with options
	comp := sc.Navbar(sc.WithID("id"), sc.WithData("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestNavbar_DefaultStyle(t *testing.T) {
	comp := sc.Navbar()
	class := comp.GetClass()
	require.Contains(t, class, "h-[50px]")
	require.Contains(t, class, "bg-gray-800")
	require.Contains(t, class, "flex")
	require.Contains(t, class, "justify-between")
	require.Contains(t, class, "text-white")
}

func TestNavbar_Output(t *testing.T) {
	id := "id"
	comp := sc.Navbar(sc.WithID(id))
	assertElementWithIDExists(t, comp, "div", id)
}

func TestNavLink_Options(t *testing.T) {
	// create navlink component with options
	comp := sc.NavLink(sc.WithID("id"), sc.WithData("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestNavLink_DefaultStyle(t *testing.T) {
	comp := sc.NavLink()
	class := comp.GetClass()
	require.Contains(t, class, "hover:bg-gray-600")
	require.Contains(t, class, "flex")
	require.Contains(t, class, "items-center")
	require.Contains(t, class, "px-2")
}

func TestNavLink_Output(t *testing.T) {
	id := "id"
	comp := sc.NavLink(sc.WithID(id))
	assertElementWithIDExists(t, comp, "a", id)
}
