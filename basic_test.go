package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	h "maragu.dev/gomponents/html"
)

func TestAny_Options(t *testing.T) {
	// create any component with options
	comp := sc.Any("div", sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
	require.Equal(t, "div", comp.GetTag())
}

func TestAny_Output(t *testing.T) {
	id := "id"

	tags := []string{"button", "a", "div", "input", "p"}
	for _, tag := range tags {
		comp := sc.Any(tag, sc.ID(id))
		assertElementWithIDExists(t, comp, tag, id)
		require.Equal(t, tag, comp.GetTag())
	}
}

func TestEmpty(t *testing.T) {
	comp := sc.Empty()
	require.Empty(t, comp.GetTag())

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Empty(t, buf.String())
}

func TestText(t *testing.T) {
	text := "This is a <div> text node!"
	escapedText := "This is a &lt;div&gt; text node!"

	comp := sc.Text(text)

	// properties
	require.Equal(t, text, comp.GetText())
	require.Empty(t, comp.GetTag())

	// content
	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, buf.String(), escapedText)
}

func TestRaw(t *testing.T) {
	text := "This is a text node!"
	comp := sc.Raw(text)

	// properties
	require.Equal(t, text, comp.GetText())
	require.Empty(t, comp.GetTag())

	// content
	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, buf.String(), text)
}

func TestWrap_Properties(t *testing.T) {
	comp := sc.Wrap(h.Div(h.ID("id")))

	require.Empty(t, comp.GetTag(), "wrap should have no tag")
	require.Empty(t, comp.GetText(), "wrap should have no text")
	require.Empty(t, comp.GetChildren(), "wrap should have no children")
	require.Empty(t, comp.GetAttrs(), "wrap should have not attributes")
}

func TestWrap_Output(t *testing.T) {
	rawNode := h.Div(h.ID("id"))
	comp := sc.Wrap(rawNode)

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	assertElementWithIDExists(t, comp, "div", "id")
}
