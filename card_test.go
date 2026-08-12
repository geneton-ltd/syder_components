package sydercomponents_test

import (
	"bytes"
	"slices"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestCard_Options(t *testing.T) {
	// create card component with options
	comp := sc.Card(sc.WithID("id"), sc.WithData("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasCardClasses(t *testing.T, class string) {
	require.Contains(t, class, "w-full")
	require.Contains(t, class, "rounded-md")
	require.Contains(t, class, "shadow-xs")
	require.Contains(t, class, "p-2")
	require.Contains(t, class, "border")
	require.Contains(t, class, "border-1")
	require.Contains(t, class, "border-gray-600")
}

func TestCard_DefaultStyle(t *testing.T) {
	comp := sc.Card()
	hasCardClasses(t, comp.GetClass())
}

func TestCard_Output(t *testing.T) {
	id := "id"
	comp := sc.Card(sc.WithID(id))

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)

	var processNode func(*html.Node) bool
	processNode = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "div" && slices.Contains(n.Attr, html.Attribute{Key: "id", Val: id}) {
			return true
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if processNode(c) {
				return true
			}
		}
		return false
	}

	node, err := html.Parse(&buf)
	require.NoError(t, err)
	require.True(t, processNode(node))
}

func TestSubCard_Merging(t *testing.T) {
	comp := sc.SubCard(sc.WithClassAppend("custom-class"))
	class := comp.GetClass()

	// has the default Card style
	hasCardClasses(t, class)

	// has its own style
	require.Contains(t, class, "border-dashed")
	require.Contains(t, class, "hover:bg-white")

	// has custom class
	require.Contains(t, class, "custom-class")
}
