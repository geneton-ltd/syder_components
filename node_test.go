package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	h "maragu.dev/gomponents/html"
)

func TestNode_Properties(t *testing.T) {
	comp := sc.Node(h.Div(h.ID("id")))

	require.Equal(t, sc.KindNode, comp.GetKind())
	require.Empty(t, comp.GetText(), "node should have no text")
	require.Empty(t, comp.GetChildren(), "node should have no children")
	require.Empty(t, comp.GetAttrs(), "node should have not attributes")
}

func TestNode_Output(t *testing.T) {
	rawNode := h.Div(h.ID("id"))
	comp := sc.Node(rawNode)

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	assertElementWithIDExists(t, comp, "div", "id")
}
