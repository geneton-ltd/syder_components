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
	require.Equal(t, text, comp.GetText())
	require.Empty(t, comp.GetTag())

	// content
	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, buf.String(), escapedText)
}

func TestRaw(t *testing.T) {
	text := "This is a <div> text node!"
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
