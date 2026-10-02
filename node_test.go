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

	require.Equal(t, sc.KindNode, string(comp.GetKind()))
	require.Empty(t, comp.GetText(), "wrap should have no text")
	require.Empty(t, comp.GetChildren(), "wrap should have no children")
	require.Empty(t, comp.GetAttrs(), "wrap should have not attributes")
}

func TestWrap_Output(t *testing.T) {
	rawNode := h.Div(h.ID("id"))
	comp := sc.Node(rawNode)

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)
	assertElementWithIDExists(t, comp, "div", "id")
}
