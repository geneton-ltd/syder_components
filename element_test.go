package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestElement_BasicMethods(t *testing.T) {
	comp := sc.Any("p", sc.Attr("key", "value"))

	require.Equal(t, sc.KindElement, comp.Kind())
	require.Equal(t, "p", comp.Name())
	require.Empty(t, comp.Value())
	require.Empty(t, comp.Children())

	v, has := comp.Attr("key")
	require.True(t, has)
	require.Equal(t, "value", v)
}

func TestElement_AttributeOverwriting(t *testing.T) {
	comp := sc.P(sc.ID("id1"), sc.ID("idfinal"))

	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "idfinal", v)
}

func TestElement_ClassMerging(t *testing.T) {
	comp := sc.P(sc.Class("font-bold"), sc.Class("flex justify-center"), sc.Class(""), sc.Class("btn"))
	v, has := comp.Attr("class")
	require.True(t, has)
	require.Equal(t, "font-bold flex justify-center btn", v)

	comp = sc.P(sc.Class(""), sc.Class("btn"))
	v, has = comp.Attr("class")
	require.True(t, has)
	require.Equal(t, "btn", v)
}

func TestElement_Children(t *testing.T) {
	component := sc.Div(
		sc.Div(),
		sc.Div(
			sc.Div(),
			sc.Div(),
		),
		sc.Div(),
	)

	// should have 3 children
	require.Len(t, component.Children(), 3)

	// 1st child has 0 children
	require.Empty(t, component.Children()[0].Children())

	// 2nd child has 2 children, each with 0 children
	require.Len(t, component.Children()[1].Children(), 2)
	require.Empty(t, component.Children()[1].Children()[0].Children())
	require.Empty(t, component.Children()[1].Children()[1].Children())

	// 3rd child has 0 children
	require.Empty(t, component.Children()[2].Children())
}

func TestElement_Render(t *testing.T) {
	t.Run("standard element", func(t *testing.T) {
		comp := sc.Span(
			sc.ID("id"),
			sc.Class("class"),
			sc.Text("text"),
		)

		var buf bytes.Buffer
		err := comp.Render(&buf)
		require.NoError(t, err)
		require.Equal(t, `<span class="class" id="id">text</span>`, buf.String())
	})

	t.Run("void elements", func(t *testing.T) {
		comp := sc.Img(sc.Src("https://images.google.com/apple"), sc.Alt("apple"))

		var buf bytes.Buffer
		err := comp.Render(&buf)
		require.NoError(t, err)
		require.Equal(t, `<img alt="apple" src="https://images.google.com/apple">`, buf.String())
	})
}

func TestElement_Unpacking(t *testing.T) {
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
