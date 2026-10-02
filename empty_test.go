package sydercomponents_test

import (
	"bytes"
	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEmpty(t *testing.T) {
	comp := sc.Empty()
	require.Equal(t, sc.KindNone, string(comp.GetKind()))

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Empty(t, buf.String())
}
