package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func TestGetters(t *testing.T) {
	simpleNavbar := sc.Navbar(
		sc.WithID("navbarID"),
		sc.WithChildren(
			sc.Any("div",
				sc.WithClass("flex"),
				sc.WithChildren(
					sc.NavLink(
						sc.WithID("link1"),
						sc.WithHref("/page1"),
						sc.WithAttributes(
							g.Text("Page1"),
						),
					),
					sc.NavLink(
						sc.WithID("link2"),
						sc.WithAttributes(
							g.Text("Page2"),
							h.Href("/page2"),
						),
					),
				),
			),
		),
	)

	assert.Equal(t, simpleNavbar.GetID(), "navbarID")
	assert.Equal(t, simpleNavbar.GetChildren()[0].GetClass(), "flex")
	h, has := simpleNavbar.GetChildren()[0].GetChildren()[1].GetHref()
	require.True(t, has)
	assert.Equal(t, h, "/page2")
}
