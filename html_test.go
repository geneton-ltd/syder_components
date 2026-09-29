package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

// <a>
func TestA_Options(t *testing.T) {
	comp := sc.A(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestA_Output(t *testing.T) {
	id := "id"
	comp := sc.A(sc.ID(id))
	assertElementWithIDExists(t, comp, "a", id)
}

// <abbr>
func TestAbbr_Options(t *testing.T) {
	comp := sc.Abbr(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestAbbr_Output(t *testing.T) {
	id := "id"
	comp := sc.Abbr(sc.ID(id))
	assertElementWithIDExists(t, comp, "abbr", id)
}

// <address>
func TestAddress_Options(t *testing.T) {
	comp := sc.Address(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestAddress_Output(t *testing.T) {
	id := "id"
	comp := sc.Address(sc.ID(id))
	assertElementWithIDExists(t, comp, "address", id)
}

// <area>
func TestArea_Options(t *testing.T) {
	comp := sc.Area(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestArea_Output(t *testing.T) {
	id := "id"
	comp := sc.Area(sc.ID(id))
	assertElementWithIDExists(t, comp, "area", id)
}

// <article>
func TestArticle_Options(t *testing.T) {
	comp := sc.Article(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestArticle_Output(t *testing.T) {
	id := "id"
	comp := sc.Article(sc.ID(id))
	assertElementWithIDExists(t, comp, "article", id)
}

// <aside>
func TestAside_Options(t *testing.T) {
	comp := sc.Aside(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestAside_Output(t *testing.T) {
	id := "id"
	comp := sc.Aside(sc.ID(id))
	assertElementWithIDExists(t, comp, "aside", id)
}

// <audio>
func TestAudio_Options(t *testing.T) {
	comp := sc.Audio(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestAudio_Output(t *testing.T) {
	id := "id"
	comp := sc.Audio(sc.ID(id))
	assertElementWithIDExists(t, comp, "audio", id)
}

// <b>
func TestB_Options(t *testing.T) {
	comp := sc.B(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestB_Output(t *testing.T) {
	id := "id"
	comp := sc.B(sc.ID(id))
	assertElementWithIDExists(t, comp, "b", id)
}

// <base>
func TestBase_Options(t *testing.T) {
	comp := sc.Base(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestBase_Output(t *testing.T) {
	id := "id"
	comp := sc.Base(sc.ID(id))
	assertElementWithIDExists(t, comp, "base", id)
}

// <bdi>
func TestBdi_Options(t *testing.T) {
	comp := sc.Bdi(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestBdi_Output(t *testing.T) {
	id := "id"
	comp := sc.Bdi(sc.ID(id))
	assertElementWithIDExists(t, comp, "bdi", id)
}

// <bdo>
func TestBdo_Options(t *testing.T) {
	comp := sc.Bdo(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestBdo_Output(t *testing.T) {
	id := "id"
	comp := sc.Bdo(sc.ID(id))
	assertElementWithIDExists(t, comp, "bdo", id)
}

// <blockquote>
func TestBlockquote_Options(t *testing.T) {
	comp := sc.Blockquote(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestBlockquote_Output(t *testing.T) {
	id := "id"
	comp := sc.Blockquote(sc.ID(id))
	assertElementWithIDExists(t, comp, "blockquote", id)
}

// <body>
func TestBody_Options(t *testing.T) {
	comp := sc.Body(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestBody_Output(t *testing.T) {
	id := "id"
	comp := sc.Body(sc.ID(id))
	assertElementWithIDExists(t, comp, "body", id)
}

// <br>
func TestBr_Options(t *testing.T) {
	comp := sc.Br(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestBr_Output(t *testing.T) {
	id := "id"
	comp := sc.Br(sc.ID(id))
	assertElementWithIDExists(t, comp, "br", id)
}

// <canvas>
func TestCanvas_Options(t *testing.T) {
	comp := sc.Canvas(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestCanvas_Output(t *testing.T) {
	id := "id"
	comp := sc.Canvas(sc.ID(id))
	assertElementWithIDExists(t, comp, "canvas", id)
}

// <cite>
func TestCite_Options(t *testing.T) {
	comp := sc.Cite(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestCite_Output(t *testing.T) {
	id := "id"
	comp := sc.Cite(sc.ID(id))
	assertElementWithIDExists(t, comp, "cite", id)
}

// <code>
func TestCode_Options(t *testing.T) {
	comp := sc.Code(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestCode_Output(t *testing.T) {
	id := "id"
	comp := sc.Code(sc.ID(id))
	assertElementWithIDExists(t, comp, "code", id)
}

// <datalist>
func TestDatalist_Options(t *testing.T) {
	comp := sc.Datalist(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestDatalist_Output(t *testing.T) {
	id := "id"
	comp := sc.Datalist(sc.ID(id))
	assertElementWithIDExists(t, comp, "datalist", id)
}

// <dd>
func TestDd_Options(t *testing.T) {
	comp := sc.Dd(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestDd_Output(t *testing.T) {
	id := "id"
	comp := sc.Dd(sc.ID(id))
	assertElementWithIDExists(t, comp, "dd", id)
}

// <del>
func TestDel_Options(t *testing.T) {
	comp := sc.Del(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestDel_Output(t *testing.T) {
	id := "id"
	comp := sc.Del(sc.ID(id))
	assertElementWithIDExists(t, comp, "del", id)
}

// <details>
func TestDetails_Options(t *testing.T) {
	comp := sc.Details(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestDetails_Output(t *testing.T) {
	id := "id"
	comp := sc.Details(sc.ID(id))
	assertElementWithIDExists(t, comp, "details", id)
}

// <dfn>
func TestDfn_Options(t *testing.T) {
	comp := sc.Dfn(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestDfn_Output(t *testing.T) {
	id := "id"
	comp := sc.Dfn(sc.ID(id))
	assertElementWithIDExists(t, comp, "dfn", id)
}

// <div>
func TestDiv_Options(t *testing.T) {
	comp := sc.Div(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestDiv_Output(t *testing.T) {
	id := "id"
	comp := sc.Div(sc.ID(id))
	assertElementWithIDExists(t, comp, "div", id)
}

// <dl>
func TestDl_Options(t *testing.T) {
	comp := sc.Dl(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestDl_Output(t *testing.T) {
	id := "id"
	comp := sc.Dl(sc.ID(id))
	assertElementWithIDExists(t, comp, "dl", id)
}

// <dt>
func TestDt_Options(t *testing.T) {
	comp := sc.Dt(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestDt_Output(t *testing.T) {
	id := "id"
	comp := sc.Dt(sc.ID(id))
	assertElementWithIDExists(t, comp, "dt", id)
}

// <em>
func TestEm_Options(t *testing.T) {
	comp := sc.Em(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestEm_Output(t *testing.T) {
	id := "id"
	comp := sc.Em(sc.ID(id))
	assertElementWithIDExists(t, comp, "em", id)
}

// <embed>
func TestEmbed_Options(t *testing.T) {
	comp := sc.Embed(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestEmbed_Output(t *testing.T) {
	id := "id"
	comp := sc.Embed(sc.ID(id))
	assertElementWithIDExists(t, comp, "embed", id)
}

// <fieldset>
func TestFieldset_Options(t *testing.T) {
	comp := sc.Fieldset(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestFieldset_Output(t *testing.T) {
	id := "id"
	comp := sc.Fieldset(sc.ID(id))
	assertElementWithIDExists(t, comp, "fieldset", id)
}

// <figcaption>
func TestFigcaption_Options(t *testing.T) {
	comp := sc.Figcaption(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestFigcaption_Output(t *testing.T) {
	id := "id"
	comp := sc.Figcaption(sc.ID(id))
	assertElementWithIDExists(t, comp, "figcaption", id)
}

// <figure>
func TestFigure_Options(t *testing.T) {
	comp := sc.Figure(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestFigure_Output(t *testing.T) {
	id := "id"
	comp := sc.Figure(sc.ID(id))
	assertElementWithIDExists(t, comp, "figure", id)
}

// <footer>
func TestFooter_Options(t *testing.T) {
	comp := sc.Footer(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestFooter_Output(t *testing.T) {
	id := "id"
	comp := sc.Footer(sc.ID(id))
	assertElementWithIDExists(t, comp, "footer", id)
}

// <form>
func TestForm_Options(t *testing.T) {
	comp := sc.Form(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestForm_Output(t *testing.T) {
	id := "id"
	comp := sc.Form(sc.ID(id))
	assertElementWithIDExists(t, comp, "form", id)
}

// <h1>
func TestH1_Options(t *testing.T) {
	comp := sc.H1(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestH1_Output(t *testing.T) {
	id := "id"
	comp := sc.H1(sc.ID(id))
	assertElementWithIDExists(t, comp, "h1", id)
}

// <h2>
func TestH2_Options(t *testing.T) {
	comp := sc.H2(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestH2_Output(t *testing.T) {
	id := "id"
	comp := sc.H2(sc.ID(id))
	assertElementWithIDExists(t, comp, "h2", id)
}

// <h3>
func TestH3_Options(t *testing.T) {
	comp := sc.H3(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestH3_Output(t *testing.T) {
	id := "id"
	comp := sc.H3(sc.ID(id))
	assertElementWithIDExists(t, comp, "h3", id)
}

// <h4>
func TestH4_Options(t *testing.T) {
	comp := sc.H4(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestH4_Output(t *testing.T) {
	id := "id"
	comp := sc.H4(sc.ID(id))
	assertElementWithIDExists(t, comp, "h4", id)
}

// <h5>
func TestH5_Options(t *testing.T) {
	comp := sc.H5(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestH5_Output(t *testing.T) {
	id := "id"
	comp := sc.H5(sc.ID(id))
	assertElementWithIDExists(t, comp, "h5", id)
}

// <h6>
func TestH6_Options(t *testing.T) {
	comp := sc.H6(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestH6_Output(t *testing.T) {
	id := "id"
	comp := sc.H6(sc.ID(id))
	assertElementWithIDExists(t, comp, "h6", id)
}

// <head>
func TestHead_Options(t *testing.T) {
	comp := sc.Head(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestHead_Output(t *testing.T) {
	id := "id"
	comp := sc.Head(sc.ID(id))
	assertElementWithIDExists(t, comp, "head", id)
}

// <header>
func TestHeader_Options(t *testing.T) {
	comp := sc.Header(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestHeader_Output(t *testing.T) {
	id := "id"
	comp := sc.Header(sc.ID(id))
	assertElementWithIDExists(t, comp, "header", id)
}

// <hgroup>
func TestHgroup_Options(t *testing.T) {
	comp := sc.Hgroup(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestHgroup_Output(t *testing.T) {
	id := "id"
	comp := sc.Hgroup(sc.ID(id))
	assertElementWithIDExists(t, comp, "hgroup", id)
}

// <hr>
func TestHr_Options(t *testing.T) {
	comp := sc.Hr(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestHr_Output(t *testing.T) {
	id := "id"
	comp := sc.Hr(sc.ID(id))
	assertElementWithIDExists(t, comp, "hr", id)
}

// <html>
func TestHtml_Options(t *testing.T) {
	comp := sc.Html(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestHtml_Output(t *testing.T) {
	id := "id"
	comp := sc.Html(sc.ID(id))
	assertElementWithIDExists(t, comp, "html", id)
}

// <i>
func TestI_Options(t *testing.T) {
	comp := sc.I(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestI_Output(t *testing.T) {
	id := "id"
	comp := sc.I(sc.ID(id))
	assertElementWithIDExists(t, comp, "i", id)
}

// <iframe>
func TestIframe_Options(t *testing.T) {
	comp := sc.Iframe(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestIframe_Output(t *testing.T) {
	id := "id"
	comp := sc.Iframe(sc.ID(id))
	assertElementWithIDExists(t, comp, "iframe", id)
}

// <img>
func TestImg_Options(t *testing.T) {
	comp := sc.Img(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestImg_Output(t *testing.T) {
	id := "id"
	comp := sc.Img(sc.ID(id))
	assertElementWithIDExists(t, comp, "img", id)
}

// <input>
func TestInput_Options(t *testing.T) {
	comp := sc.Input(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestInput_Output(t *testing.T) {
	id := "id"
	comp := sc.Input(sc.ID(id))
	assertElementWithIDExists(t, comp, "input", id)
}

// <ins>
func TestIns_Options(t *testing.T) {
	comp := sc.Ins(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestIns_Output(t *testing.T) {
	id := "id"
	comp := sc.Ins(sc.ID(id))
	assertElementWithIDExists(t, comp, "ins", id)
}

// <kbd>
func TestKbd_Options(t *testing.T) {
	comp := sc.Kbd(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestKbd_Output(t *testing.T) {
	id := "id"
	comp := sc.Kbd(sc.ID(id))
	assertElementWithIDExists(t, comp, "kbd", id)
}

// <label>
func TestLabel_Options(t *testing.T) {
	comp := sc.Label(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestLabel_Output(t *testing.T) {
	id := "id"
	comp := sc.Label(sc.ID(id))
	assertElementWithIDExists(t, comp, "label", id)
}

// <legend>
func TestLegend_Options(t *testing.T) {
	comp := sc.Legend(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestLegend_Output(t *testing.T) {
	id := "id"
	comp := sc.Legend(sc.ID(id))
	assertElementWithIDExists(t, comp, "legend", id)
}

// <li>
func TestLi_Options(t *testing.T) {
	comp := sc.Li(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestLi_Output(t *testing.T) {
	id := "id"
	comp := sc.Li(sc.ID(id))
	assertElementWithIDExists(t, comp, "li", id)
}

// <link>
func TestLink_Options(t *testing.T) {
	comp := sc.Link(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestLink_Output(t *testing.T) {
	id := "id"
	comp := sc.Link(sc.ID(id))
	assertElementWithIDExists(t, comp, "link", id)
}

// <main>
func TestMain_Options(t *testing.T) {
	comp := sc.Main(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestMain_Output(t *testing.T) {
	id := "id"
	comp := sc.Main(sc.ID(id))
	assertElementWithIDExists(t, comp, "main", id)
}

// <map>
func TestMap_Options(t *testing.T) {
	comp := sc.Map(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestMap_Output(t *testing.T) {
	id := "id"
	comp := sc.Map(sc.ID(id))
	assertElementWithIDExists(t, comp, "map", id)
}

// <mark>
func TestMark_Options(t *testing.T) {
	comp := sc.Mark(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestMark_Output(t *testing.T) {
	id := "id"
	comp := sc.Mark(sc.ID(id))
	assertElementWithIDExists(t, comp, "mark", id)
}

// <menu>
func TestMenu_Options(t *testing.T) {
	comp := sc.Menu(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestMenu_Output(t *testing.T) {
	id := "id"
	comp := sc.Menu(sc.ID(id))
	assertElementWithIDExists(t, comp, "menu", id)
}

// <meta>
func TestMeta_Options(t *testing.T) {
	comp := sc.Meta(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestMeta_Output(t *testing.T) {
	id := "id"
	comp := sc.Meta(sc.ID(id))
	assertElementWithIDExists(t, comp, "meta", id)
}

// <meter>
func TestMeter_Options(t *testing.T) {
	comp := sc.Meter(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestMeter_Output(t *testing.T) {
	id := "id"
	comp := sc.Meter(sc.ID(id))
	assertElementWithIDExists(t, comp, "meter", id)
}

// <nav>
func TestNav_Options(t *testing.T) {
	comp := sc.Nav(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestNav_Output(t *testing.T) {
	id := "id"
	comp := sc.Nav(sc.ID(id))
	assertElementWithIDExists(t, comp, "nav", id)
}

// <noscript>
func TestNoscript_Options(t *testing.T) {
	comp := sc.Noscript(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestNoscript_Output(t *testing.T) {
	id := "id"
	comp := sc.Noscript(sc.ID(id))
	assertElementWithIDExists(t, comp, "noscript", id)
}

// <object>
func TestObject_Options(t *testing.T) {
	comp := sc.Object(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestObject_Output(t *testing.T) {
	id := "id"
	comp := sc.Object(sc.ID(id))
	assertElementWithIDExists(t, comp, "object", id)
}

// <ol>
func TestOl_Options(t *testing.T) {
	comp := sc.Ol(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestOl_Output(t *testing.T) {
	id := "id"
	comp := sc.Ol(sc.ID(id))
	assertElementWithIDExists(t, comp, "ol", id)
}

// <optgroup>
func TestOptgroup_Options(t *testing.T) {
	comp := sc.Optgroup(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestOptgroup_Output(t *testing.T) {
	id := "id"
	comp := sc.Optgroup(sc.ID(id))
	assertElementWithIDExists(t, comp, "optgroup", id)
}

// <output>
func TestOutput_Options(t *testing.T) {
	comp := sc.Output(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestOutput_Output(t *testing.T) {
	id := "id"
	comp := sc.Output(sc.ID(id))
	assertElementWithIDExists(t, comp, "output", id)
}

// <p>
func TestP_Options(t *testing.T) {
	comp := sc.P(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestP_Output(t *testing.T) {
	id := "id"
	comp := sc.P(sc.ID(id))
	assertElementWithIDExists(t, comp, "p", id)
}

// <picture>
func TestPicture_Options(t *testing.T) {
	comp := sc.Picture(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestPicture_Output(t *testing.T) {
	id := "id"
	comp := sc.Picture(sc.ID(id))
	assertElementWithIDExists(t, comp, "picture", id)
}

// <pre>
func TestPre_Options(t *testing.T) {
	comp := sc.Pre(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestPre_Output(t *testing.T) {
	id := "id"
	comp := sc.Pre(sc.ID(id))
	assertElementWithIDExists(t, comp, "pre", id)
}

// <progress>
func TestProgress_Options(t *testing.T) {
	comp := sc.Progress(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestProgress_Output(t *testing.T) {
	id := "id"
	comp := sc.Progress(sc.ID(id))
	assertElementWithIDExists(t, comp, "progress", id)
}

// <q>
func TestQ_Options(t *testing.T) {
	comp := sc.Q(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestQ_Output(t *testing.T) {
	id := "id"
	comp := sc.Q(sc.ID(id))
	assertElementWithIDExists(t, comp, "q", id)
}

// <rp>
func TestRp_Options(t *testing.T) {
	comp := sc.Rp(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestRp_Output(t *testing.T) {
	id := "id"
	comp := sc.Rp(sc.ID(id))
	assertElementWithIDExists(t, comp, "rp", id)
}

// <rt>
func TestRt_Options(t *testing.T) {
	comp := sc.Rt(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestRt_Output(t *testing.T) {
	id := "id"
	comp := sc.Rt(sc.ID(id))
	assertElementWithIDExists(t, comp, "rt", id)
}

// <ruby>
func TestRuby_Options(t *testing.T) {
	comp := sc.Ruby(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestRuby_Output(t *testing.T) {
	id := "id"
	comp := sc.Ruby(sc.ID(id))
	assertElementWithIDExists(t, comp, "ruby", id)
}

// <s>
func TestS_Options(t *testing.T) {
	comp := sc.S(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestS_Output(t *testing.T) {
	id := "id"
	comp := sc.S(sc.ID(id))
	assertElementWithIDExists(t, comp, "s", id)
}

// <samp>
func TestSamp_Options(t *testing.T) {
	comp := sc.Samp(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSamp_Output(t *testing.T) {
	id := "id"
	comp := sc.Samp(sc.ID(id))
	assertElementWithIDExists(t, comp, "samp", id)
}

// <script>
func TestScript_Options(t *testing.T) {
	comp := sc.Script(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestScript_Output(t *testing.T) {
	id := "id"
	comp := sc.Script(sc.ID(id))
	assertElementWithIDExists(t, comp, "script", id)
}

// <search>
func TestSearch_Options(t *testing.T) {
	comp := sc.Search(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSearch_Output(t *testing.T) {
	id := "id"
	comp := sc.Search(sc.ID(id))
	assertElementWithIDExists(t, comp, "search", id)
}

// <section>
func TestSection_Options(t *testing.T) {
	comp := sc.Section(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSection_Output(t *testing.T) {
	id := "id"
	comp := sc.Section(sc.ID(id))
	assertElementWithIDExists(t, comp, "section", id)
}

// <select>
func TestSelect_Options(t *testing.T) {
	comp := sc.Select(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSelect_Output(t *testing.T) {
	id := "id"
	comp := sc.Select(sc.ID(id))
	assertElementWithIDExists(t, comp, "select", id)
}

// <selectedcontent>
func TestSelectedcontent_Options(t *testing.T) {
	comp := sc.Selectedcontent(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSelectedcontent_Output(t *testing.T) {
	id := "id"
	comp := sc.Selectedcontent(sc.ID(id))
	assertElementWithIDExists(t, comp, "selectedcontent", id)
}

// <slot>
func TestSlot_Options(t *testing.T) {
	comp := sc.Slot(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSlot_Output(t *testing.T) {
	id := "id"
	comp := sc.Slot(sc.ID(id))
	assertElementWithIDExists(t, comp, "slot", id)
}

// <small>
func TestSmall_Options(t *testing.T) {
	comp := sc.Small(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSmall_Output(t *testing.T) {
	id := "id"
	comp := sc.Small(sc.ID(id))
	assertElementWithIDExists(t, comp, "small", id)
}

// <source>
func TestSource_Options(t *testing.T) {
	comp := sc.Source(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSource_Output(t *testing.T) {
	id := "id"
	comp := sc.Source(sc.ID(id))
	assertElementWithIDExists(t, comp, "source", id)
}

// <span>
func TestSpan_Options(t *testing.T) {
	comp := sc.Span(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSpan_Output(t *testing.T) {
	id := "id"
	comp := sc.Span(sc.ID(id))
	assertElementWithIDExists(t, comp, "span", id)
}

// <strong>
func TestStrong_Options(t *testing.T) {
	comp := sc.Strong(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestStrong_Output(t *testing.T) {
	id := "id"
	comp := sc.Strong(sc.ID(id))
	assertElementWithIDExists(t, comp, "strong", id)
}

// <sub>
func TestSub_Options(t *testing.T) {
	comp := sc.Sub(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSub_Output(t *testing.T) {
	id := "id"
	comp := sc.Sub(sc.ID(id))
	assertElementWithIDExists(t, comp, "sub", id)
}

// <summary>
func TestSummary_Options(t *testing.T) {
	comp := sc.Summary(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSummary_Output(t *testing.T) {
	id := "id"
	comp := sc.Summary(sc.ID(id))
	assertElementWithIDExists(t, comp, "summary", id)
}

// <sup>
func TestSup_Options(t *testing.T) {
	comp := sc.Sup(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestSup_Output(t *testing.T) {
	id := "id"
	comp := sc.Sup(sc.ID(id))
	assertElementWithIDExists(t, comp, "sup", id)
}

// <template>
func TestTemplate_Options(t *testing.T) {
	comp := sc.Template(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestTemplate_Output(t *testing.T) {
	id := "id"
	comp := sc.Template(sc.ID(id))
	assertElementWithIDExists(t, comp, "template", id)
}

// <textarea>
func TestTextarea_Options(t *testing.T) {
	comp := sc.Textarea(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestTextarea_Output(t *testing.T) {
	id := "id"
	comp := sc.Textarea(sc.ID(id))
	assertElementWithIDExists(t, comp, "textarea", id)
}

// <time>
func TestTime_Options(t *testing.T) {
	comp := sc.Time(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestTime_Output(t *testing.T) {
	id := "id"
	comp := sc.Time(sc.ID(id))
	assertElementWithIDExists(t, comp, "time", id)
}

// <title>
func TestTitle_Options(t *testing.T) {
	comp := sc.Title(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestTitle_Output(t *testing.T) {
	id := "id"
	comp := sc.Title(sc.ID(id))
	assertElementWithIDExists(t, comp, "title", id)
}

// <track>
func TestTrack_Options(t *testing.T) {
	comp := sc.Track(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestTrack_Output(t *testing.T) {
	id := "id"
	comp := sc.Track(sc.ID(id))
	assertElementWithIDExists(t, comp, "track", id)
}

// <u>
func TestU_Options(t *testing.T) {
	comp := sc.U(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestU_Output(t *testing.T) {
	id := "id"
	comp := sc.U(sc.ID(id))
	assertElementWithIDExists(t, comp, "u", id)
}

// <ul>
func TestUl_Options(t *testing.T) {
	comp := sc.Ul(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestUl_Output(t *testing.T) {
	id := "id"
	comp := sc.Ul(sc.ID(id))
	assertElementWithIDExists(t, comp, "ul", id)
}

// <var>
func TestVar_Options(t *testing.T) {
	comp := sc.Var(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestVar_Output(t *testing.T) {
	id := "id"
	comp := sc.Var(sc.ID(id))
	assertElementWithIDExists(t, comp, "var", id)
}

// <video>
func TestVideo_Options(t *testing.T) {
	comp := sc.Video(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestVideo_Output(t *testing.T) {
	id := "id"
	comp := sc.Video(sc.ID(id))
	assertElementWithIDExists(t, comp, "video", id)
}

// <wbr>
func TestWbr_Options(t *testing.T) {
	comp := sc.Wbr(sc.ID("id"), sc.Data("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestWbr_Output(t *testing.T) {
	id := "id"
	comp := sc.Wbr(sc.ID(id))
	assertElementWithIDExists(t, comp, "wbr", id)
}
