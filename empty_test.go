package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func EmptyTest(t *testing.T) {
	comp := sc.Empty()

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Empty(t, buf.String())
}
