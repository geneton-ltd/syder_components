package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestTextComponent(t *testing.T) {
	text := "This is a text node!"
	comp := sc.Text(text)

	require.Equal(t, text, comp.GetText())

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, buf.String(), text)
}
