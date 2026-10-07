package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestGroup(t *testing.T) {
	nodes := []sc.Node{
		sc.Div(),
		sc.A(sc.Href("google.com"), sc.Text("Click Me!")),
		sc.Style("padding: 0"),
	}
	g := sc.Group(nodes...)

	require.Equal(t, sc.KindGroup, g.Kind())
	require.Empty(t, g.Name())
	require.Empty(t, g.Value())
	require.Equal(t, nodes, g.Children())

	var buf bytes.Buffer
	err := g.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, `<div></div><a href="google.com">Click Me!</a>style="padding: 0"`, buf.String())
}

func TestGroup_Empty(t *testing.T) {
	g := sc.Group()
	require.Equal(t, sc.KindGroup, g.Kind())
	require.Empty(t, g.Name())
	require.Empty(t, g.Value())
	require.Empty(t, g.Children())

	var buf bytes.Buffer
	err := g.Render(&buf)
	require.NoError(t, err)
	require.Empty(t, buf.String())
}

func TestGroup_UnpackingInElement(t *testing.T) {
	group := sc.Group(
		sc.Class("items-center justify-center"),
		sc.ID("id123"),
		sc.Group(sc.Type("submit"), sc.Alt("image")),
		sc.Span(sc.Text("some content")),
		sc.Group(sc.Div(sc.ID("1")), sc.Div(sc.ID("2"))),
	)

	div := sc.Div(group)

	// has attributes
	class, hasClass := div.Attr("class")
	require.True(t, hasClass)
	require.Equal(t, "items-center justify-center", class)

	id, hasID := div.Attr("id")
	require.True(t, hasID)
	require.Equal(t, "id123", id)

	typeAttr, hasTypeAttr := div.Attr("type")
	require.True(t, hasTypeAttr)
	require.Equal(t, "submit", typeAttr)

	alt, hasAlt := div.Attr("alt")
	require.True(t, hasAlt)
	require.Equal(t, "image", alt)

	// has children
	children := div.Children()
	require.Len(t, children, 3)
	require.Equal(t, "span", children[0].Name())
	require.Equal(t, "div", children[1].Name())
	require.Equal(t, "div", children[2].Name())

	// check rendered content
	var buf bytes.Buffer
	err := div.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, `<div alt="image" class="items-center justify-center" id="id123" type="submit"><span>some content</span><div id="1"></div><div id="2"></div></div>`, buf.String())
}
