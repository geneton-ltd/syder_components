package sydercomponents_test

import (
	"bytes"
	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmpty(t *testing.T) {
	comp := sc.Empty()
	require.Empty(t, comp.GetTag())

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Empty(t, buf.String())
}
