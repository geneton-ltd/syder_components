package sydercomponents_test

import (
	"bytes"
	"slices"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func assertElementWithIDExists(t *testing.T, comp sc.Component, elementTag string) {
	id := comp.GetID()

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)

	var processNode func(*html.Node) bool
	processNode = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == elementTag && slices.Contains(n.Attr, html.Attribute{Key: "id", Val: id}) {
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
