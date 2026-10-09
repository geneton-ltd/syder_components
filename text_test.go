package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestText(t *testing.T) {
	text := "This is a <div> text node!"
	escapedText := "This is a &lt;div&gt; text node!"

	comp := sc.Text(text)

	// properties
	require.Equal(t, sc.KindText, comp.Kind())
	require.Empty(t, comp.Name())
	require.Equal(t, escapedText, comp.Value())
	require.Empty(t, comp.Children())

	// content
	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, escapedText, buf.String())
}

func TestRaw(t *testing.T) {
	text := "This is a <div> text node!"
	comp := sc.Raw(text)

	// properties
	require.Equal(t, sc.KindText, comp.Kind())
	require.Empty(t, comp.Name())
	require.Equal(t, text, comp.Value())
	require.Empty(t, comp.Children())

	// content
	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, text, buf.String())
}
