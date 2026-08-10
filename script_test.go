package sydercomponents_test

import (
	"bytes"
	"slices"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func assertScriptNodeExists(t *testing.T, node *html.Node, scriptId, content string) {
	found := false

	var processNode func(*html.Node)
	processNode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" && slices.Contains(n.Attr, html.Attribute{Key: "id", Val: scriptId}) {
			require.NotNil(t, n.FirstChild)
			require.Equal(t, html.TextNode, n.FirstChild.Type)
			require.Equal(t, content, n.FirstChild.Data)
			found = true
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			processNode(c)
		}
	}
	processNode(node)
	require.True(t, found)
}

func TestScript_WithRegistry(t *testing.T) {
	ctx := t.Context()
	reg := &sc.ScriptRegistry{}
	ctx = sc.WithScriptRegistry(ctx, reg)

	content := "const hello = new Car(); console.log(hello);"
	scriptId := "id"
	comp := sc.Script(ctx, content, sc.WithID(scriptId))

	const contentHash = "'sha512-ODmofMDdZV3dCsE94FZUWHwtMouk7BzHXWAjVOFqrtoSst93EDy1wbsty5NrYT1OlJi6BRJKpfcy0CY2gt8gTA=='"
	require.Len(t, reg.Hashes, 1)
	require.Equal(t, contentHash, reg.Hashes[0])

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)

	node, err := html.Parse(&buf)
	require.NoError(t, err)
	assertScriptNodeExists(t, node, scriptId, content)
}

func TestScript_WithoutRegistry(t *testing.T) {
	ctx := t.Context()

	content := "const hello = new Car(); console.log(hello);"
	scriptId := "id"
	comp := sc.Script(ctx, content, sc.WithID(scriptId))

	var buf bytes.Buffer
	err := comp.Render(&buf)
	require.NoError(t, err)

	node, err := html.Parse(&buf)
	require.NoError(t, err)
	assertScriptNodeExists(t, node, scriptId, content)
}
