package sydercomponents_test

import (
	"bytes"
	"fmt"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestHTMLElements_Basic(t *testing.T) {
	testcases := []struct {
		tag  string
		comp func(children ...sc.Node) sc.Node
	}{
		{tag: "a", comp: sc.A},
		{tag: "abbr", comp: sc.Abbr},
		{tag: "address", comp: sc.Address},
		{tag: "area", comp: sc.Area},
		{tag: "article", comp: sc.Article},
		{tag: "aside", comp: sc.Aside},
		{tag: "audio", comp: sc.Audio},
		{tag: "b", comp: sc.B},
		{tag: "base", comp: sc.Base},
		{tag: "bdi", comp: sc.Bdi},
		{tag: "bdo", comp: sc.Bdo},
		{tag: "blockquote", comp: sc.Blockquote},
		{tag: "body", comp: sc.Body},
		{tag: "br", comp: sc.Br},
		{tag: "button", comp: sc.Button},
		{tag: "canvas", comp: sc.Canvas},
		{tag: "cite", comp: sc.Cite},
		{tag: "code", comp: sc.Code},
		{tag: "data", comp: sc.DataEl},
		{tag: "datalist", comp: sc.Datalist},
		{tag: "dd", comp: sc.Dd},
		{tag: "del", comp: sc.Del},
		{tag: "details", comp: sc.Details},
		{tag: "dfn", comp: sc.Dfn},
		{tag: "dialog", comp: sc.Dialog},
		{tag: "div", comp: sc.Div},
		{tag: "dl", comp: sc.Dl},
		{tag: "dt", comp: sc.Dt},
		{tag: "em", comp: sc.Em},
		{tag: "embed", comp: sc.Embed},
		{tag: "fieldset", comp: sc.Fieldset},
		{tag: "figcaption", comp: sc.Figcaption},
		{tag: "figure", comp: sc.Figure},
		{tag: "footer", comp: sc.Footer},
		{tag: "form", comp: sc.Form},
		{tag: "h1", comp: sc.H1},
		{tag: "h2", comp: sc.H2},
		{tag: "h3", comp: sc.H3},
		{tag: "h4", comp: sc.H4},
		{tag: "h5", comp: sc.H5},
		{tag: "h6", comp: sc.H6},
		{tag: "head", comp: sc.Head},
		{tag: "header", comp: sc.Header},
		{tag: "hgroup", comp: sc.Hgroup},
		{tag: "hr", comp: sc.Hr},
		{tag: "html", comp: sc.Html},
		{tag: "i", comp: sc.I},
		{tag: "iframe", comp: sc.Iframe},
		{tag: "img", comp: sc.Img},
		{tag: "input", comp: sc.Input},
		{tag: "ins", comp: sc.Ins},
		{tag: "kbd", comp: sc.Kbd},
		{tag: "label", comp: sc.Label},
		{tag: "legend", comp: sc.Legend},
		{tag: "li", comp: sc.Li},
		{tag: "link", comp: sc.Link},
		{tag: "main", comp: sc.Main},
		{tag: "map", comp: sc.Map},
		{tag: "mark", comp: sc.Mark},
		{tag: "menu", comp: sc.Menu},
		{tag: "meta", comp: sc.Meta},
		{tag: "meter", comp: sc.Meter},
		{tag: "nav", comp: sc.Nav},
		{tag: "noscript", comp: sc.Noscript},
		{tag: "object", comp: sc.Object},
		{tag: "ol", comp: sc.Ol},
		{tag: "optgroup", comp: sc.Optgroup},
		{tag: "option", comp: sc.Option},
		{tag: "output", comp: sc.Output},
		{tag: "p", comp: sc.P},
		{tag: "picture", comp: sc.Picture},
		{tag: "pre", comp: sc.Pre},
		{tag: "progress", comp: sc.Progress},
		{tag: "q", comp: sc.Q},
		{tag: "rp", comp: sc.Rp},
		{tag: "rt", comp: sc.Rt},
		{tag: "ruby", comp: sc.Ruby},
		{tag: "s", comp: sc.S},
		{tag: "samp", comp: sc.Samp},
		{tag: "script", comp: sc.Script},
		{tag: "search", comp: sc.Search},
		{tag: "section", comp: sc.Section},
		{tag: "select", comp: sc.Select},
		{tag: "selectedcontent", comp: sc.Selectedcontent},
		{tag: "slot", comp: sc.Slot},
		{tag: "small", comp: sc.Small},
		{tag: "source", comp: sc.Source},
		{tag: "span", comp: sc.Span},
		{tag: "strong", comp: sc.Strong},
		{tag: "style", comp: sc.StyleEl},
		{tag: "sub", comp: sc.Sub},
		{tag: "summary", comp: sc.Summary},
		{tag: "sup", comp: sc.Sup},
		{tag: "table", comp: sc.Table},
		{tag: "textarea", comp: sc.Textarea},
		{tag: "time", comp: sc.Time},
		{tag: "title", comp: sc.Title},
		{tag: "track", comp: sc.Track},
		{tag: "u", comp: sc.U},
		{tag: "ul", comp: sc.Ul},
		{tag: "var", comp: sc.Var},
		{tag: "video", comp: sc.Video},
		{tag: "wbr", comp: sc.Wbr},
	}

	for _, tt := range testcases {
		t.Run(tt.tag+" options", func(t *testing.T) {
			comp := tt.comp(sc.ID("id123"), sc.Data("name", "john"))

			id, has := comp.Attr("id")
			require.Equal(t, "id123", id)
			require.True(t, has)
			v, has := comp.Attr("data-name")
			require.Equal(t, "john", v)
			require.True(t, has)
		})

		t.Run(tt.tag+" output", func(t *testing.T) {
			id := "id123"
			comp := tt.comp(sc.ID(id))
			tu.AssertElementWithIDExists(t, comp, tt.tag, id)
		})
	}
}

func TestHTMLElements_Table(t *testing.T) {
	testcases := []struct {
		tag         string
		comp        func(children ...sc.Node) sc.Node
		selfClosing bool
	}{
		{tag: "caption", comp: sc.Caption},
		{tag: "col", comp: sc.Col, selfClosing: true},
		{tag: "colgroup", comp: sc.Colgroup},
		{tag: "tbody", comp: sc.TBody},
		{tag: "td", comp: sc.Td},
		{tag: "tfoot", comp: sc.Tfoot},
		{tag: "th", comp: sc.Th},
		{tag: "thead", comp: sc.THead},
		{tag: "tr", comp: sc.Tr},
	}

	for _, tt := range testcases {
		comp := sc.Table(tt.comp())
		tagContent := fmt.Sprintf("<%s></%s>", tt.tag, tt.tag)
		if tt.selfClosing {
			tagContent = fmt.Sprintf("<%s>", tt.tag)
		}

		var buf bytes.Buffer
		err := comp.Render(&buf)
		require.NoError(t, err)
		require.Equal(t, fmt.Sprintf("<table>%s</table>", tagContent), buf.String())
	}
}
