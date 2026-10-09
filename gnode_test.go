package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	"maragu.dev/gomponents-heroicons/v3/solid"
)

func TestGNode(t *testing.T) {
	comp := sc.GNode(solid.ArrowLeft())

	require.Equal(t, sc.KindGNode, comp.Kind())
	require.Empty(t, comp.Name())
	require.Empty(t, comp.Value())
	require.Empty(t, comp.Children())

	v, has := comp.Attr("any")
	require.False(t, has)
	require.Empty(t, v)

	// output is should be same as gomponents.Render
	var bufComp bytes.Buffer
	err := comp.Render(&bufComp)
	require.NoError(t, err)

	var bufGNode bytes.Buffer
	err = solid.ArrowLeft().Render(&bufGNode)
	require.NoError(t, err)
	require.Equal(t, bufGNode.String(), bufComp.String())
}

func TestGNode_Empty(t *testing.T) {
	comp := sc.GNode(nil)

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	require.Empty(t, buf.String())
}
