package sydercomponents_test

import (
	"bytes"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestGroup(t *testing.T) {
	nodes := []sc.Node{
		sc.Div(),
		sc.A(sc.Href("google.com"), sc.Text("Click Me!")),
		sc.Style("padding: 0"),
	}
	g := sc.Group(nodes...)

	require.Equal(t, sc.KindGroup, g.Kind())
	require.Empty(t, g.Name())
	require.Empty(t, g.Value())
	require.Equal(t, nodes, g.Children())

	var buf bytes.Buffer
	err := g.Render(&buf)
	require.NoError(t, err)
	require.Equal(t, `<div></div><a href="google.com">Click Me!</a>style="padding: 0"`, buf.String())
}

func TestGroup_Empty(t *testing.T) {
	g := sc.Group()
	require.Equal(t, sc.KindGroup, g.Kind())
	require.Empty(t, g.Name())
	require.Empty(t, g.Value())
	require.Empty(t, g.Children())

	var buf bytes.Buffer
	err := g.Render(&buf)
	require.NoError(t, err)
	require.Empty(t, buf.String())
}
