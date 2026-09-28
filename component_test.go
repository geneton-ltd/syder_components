package sydercomponents_test

import (
	"strings"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	sydercomponents "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	g "maragu.dev/gomponents"
)

func TestComponent_SimpleOptions(t *testing.T) {
	testcases := []struct {
		name  string
		comp  sc.Component
		check func(t *testing.T, comp sc.Component)
	}{
		{
			name: "test WithNode",
			comp: sc.Any("div", sc.Nodes(g.Text("text"))),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, []g.Node{g.Text("text")}, comp.GetNodes())
			},
		},
		{
			name: "test WithNode overwrites",
			comp: sc.Any("div", sc.Nodes(g.Text("text")), sc.Nodes(g.Text("text2"))),
			check: func(t *testing.T, comp sc.Component) {
				require.Contains(t, comp.GetNodes(), g.Text("text2"))
				require.NotContains(t, comp.GetNodes(), g.Text("text"))
			},
		},
		{
			name: "test WithNodeAppend",
			comp: sc.Any("div", sc.Nodes(g.Text("text")), sc.NodeAppend(g.Text("text2"))),
			check: func(t *testing.T, comp sc.Component) {
				require.Contains(t, comp.GetNodes(), g.Text("text"))
				require.Contains(t, comp.GetNodes(), g.Text("text2"))
			},
		},
		{
			name: "test WithClass",
			comp: sc.Any("div", sc.Class("flex")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "flex", comp.GetClass())
			},
		},
		{
			name: "test WithClass overwrites",
			comp: sc.Any("div", sc.Class("flex"), sc.Class("none")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "none", comp.GetClass())
				require.NotEqual(t, "flex", comp.GetClass())
			},
		},
		{
			name: "test WithId",
			comp: sc.Any("div", sc.ID("id")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "id", comp.GetID())
			},
		},
		{
			name: "test WithData",
			comp: sc.Any("div", sc.Data("apple", "1"), sc.Data("token", "secret")),
			check: func(t *testing.T, comp sc.Component) {
				v, has := comp.GetData("apple")
				require.Equal(t, v, "1")
				require.True(t, has)

				v, has = comp.GetData("token")
				require.Equal(t, v, "secret")
				require.True(t, has)

				v, has = comp.GetData("INVALID")
				require.Empty(t, v)
				require.False(t, has)
			},
		},
		{
			name: "test WithAria",
			comp: sc.Any("div", sc.Aria("apple", "1"), sc.Aria("token", "secret")),
			check: func(t *testing.T, comp sc.Component) {
				v, has := comp.GetAria("apple")
				require.Equal(t, v, "1")
				require.True(t, has)

				v, has = comp.GetAria("token")
				require.Equal(t, v, "secret")
				require.True(t, has)

				v, has = comp.GetAria("INVALID")
				require.Empty(t, v)
				require.False(t, has)
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.check(t, testcase.comp)
		})
	}
}

func TestComponent_OptionalStringOptions(t *testing.T) {
	testcases := []struct {
		name      string
		attrValue string
		optFn     func(string) sc.Option
		getter    func(sydercomponents.Component) (string, bool)
	}{
		{
			name:      "test style",
			attrValue: "background-color: black",
			optFn:     sc.Style,
			getter:    sc.Component.GetStyle,
		},
		{
			name:      "test name",
			attrValue: "metadataNameButton",
			optFn:     sc.Name,
			getter:    sc.Component.GetName,
		},
		{
			name:      "test type",
			attrValue: "button",
			optFn:     sc.Type,
			getter:    sc.Component.GetType,
		},
		{
			name:      "test placeholder",
			attrValue: "Enter your CCV from the back of your card...",
			optFn:     sc.Placeholder,
			getter:    sc.Component.GetPlaceholder,
		},
		{
			name:      "test value",
			attrValue: "67",
			optFn:     sc.Value,
			getter:    sc.Component.GetValue,
		},
		{
			name:      "test src",
			attrValue: "screenshotted_nft.jpg",
			optFn:     sc.Src,
			getter:    sc.Component.GetSrc,
		},
		{
			name:      "test alt",
			attrValue: "person with a cat",
			optFn:     sc.Alt,
			getter:    sc.Component.GetAlt,
		},
		{
			name:      "test href",
			attrValue: "http://globalhost",
			optFn:     sc.Href,
			getter:    sc.Component.GetHref,
		},
		{
			name:      "test target",
			attrValue: "_blank",
			optFn:     sc.Target,
			getter:    sc.Component.GetTarget,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			comp := sc.Any("div", testcase.optFn(testcase.attrValue))
			v, has := testcase.getter(comp)
			require.Equal(t, testcase.attrValue, v)
			require.True(t, has)

			comp = sc.Any("div")
			v, has = testcase.getter(comp)
			require.Empty(t, v)
			require.False(t, has)
		})
	}
}

func TestComponent_EdgeCaseOptions(t *testing.T) {
	testcases := []struct {
		name  string
		comp  sc.Component
		check func(t *testing.T, comp sc.Component)
	}{
		{
			name: "test WithTabIndex",
			comp: sc.Any("div", sc.TabIndex(-1)),
			check: func(t *testing.T, comp sc.Component) {
				v, has := comp.GetTabIndex()
				require.Equal(t, -1, v)
				require.True(t, has)
			},
		},
		{
			name: "test WithTabIndex missing",
			comp: sc.Any("div"),
			check: func(t *testing.T, comp sc.Component) {
				v, has := comp.GetTabIndex()
				require.Empty(t, v)
				require.False(t, has)
			},
		},
		{
			name: "test WithClassAppend",
			comp: sc.Any("div", sc.ClassAppend("flex"), sc.ClassAppend("flex-col")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "flex flex-col", comp.GetClass())
			},
		},
		{
			name: "test WithClass and WithClassAppend",
			comp: sc.Any("div", sc.Class("flex"), sc.ClassAppend("flex-col")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "flex flex-col", comp.GetClass())
			},
		},
		{
			name: "test WithClassAppend on empty class",
			comp: sc.Any("div", sc.ClassAppend("  flex "), sc.ClassAppend(" flex-col")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "flex flex-col", comp.GetClass())
			},
		},
		{
			name: "test WithDisabled",
			comp: sc.Any("div", sc.Disabled()),
			check: func(t *testing.T, comp sc.Component) {
				require.True(t, comp.HasDisabled())
			},
		},
		{
			name: "test WithDisabled missing",
			comp: sc.Any("div"),
			check: func(t *testing.T, comp sc.Component) {
				require.False(t, comp.HasDisabled())
			},
		},
		{
			name: "test WithRequired",
			comp: sc.Any("div", sc.Required()),
			check: func(t *testing.T, comp sc.Component) {
				require.True(t, comp.HasRequired())
			},
		},
		{
			name: "test WithRequired missing",
			comp: sc.Any("div"),
			check: func(t *testing.T, comp sc.Component) {
				require.False(t, comp.HasRequired())
			},
		},
		{
			name: "test WithHidden",
			comp: sc.Any("div", sc.Hidden()),
			check: func(t *testing.T, comp sc.Component) {
				require.True(t, comp.HasHidden())
			},
		},
		{
			name: "test WithHidden missing",
			comp: sc.Any("div"),
			check: func(t *testing.T, comp sc.Component) {
				require.False(t, comp.HasHidden())
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.check(t, testcase.comp)
		})
	}
}

func TestComponent_Children(t *testing.T) {
	t.Run("GetChildren", func(t *testing.T) {
		component := sc.Any("div",
			sc.Children(
				sc.Any("div"),
				sc.Any("div",
					sc.Children(
						sc.Any("div"),
						sc.Any("div"),
					),
				),
				sc.Any("div"),
			),
		)

		// should have 3 children
		require.Len(t, component.GetChildren(), 3)

		// 1st child has 0 children
		require.Empty(t, component.GetChildren()[0].GetChildren())

		// 2nd child has 2 children, each with 0 children
		require.Len(t, component.GetChildren()[1].GetChildren(), 2)
		require.Empty(t, component.GetChildren()[1].GetChildren()[0].GetChildren())
		require.Empty(t, component.GetChildren()[1].GetChildren()[1].GetChildren())

		// 3rd child has 0 children
		require.Empty(t, component.GetChildren()[2].GetChildren())
	})

	t.Run("WithChildren overwrite", func(t *testing.T) {
		component := sc.Any("div",
			sc.Children(sc.Any("div"), sc.Any("div")),
			sc.Children(sc.Any("div")),
		)

		require.Len(t, component.GetChildren(), 1, "should have 1 child on overwrite")
	})

	t.Run("WithChildren and WithChildrenAppend", func(t *testing.T) {
		component := sc.Any("div",
			sc.Children(sc.Any("div"), sc.Any("div")),
			sc.ChildrenAppend(sc.Any("div")),
		)

		require.Len(t, component.GetChildren(), 3, "should have 3 children with append")
	})
}

func TestComponent_Render(t *testing.T) {
	comp := sc.Any("span",
		sc.ID("id"),
		sc.Class("class"),
		sc.Nodes(g.Text("text")),
	)

	var sb strings.Builder
	err := comp.Render(&sb)
	require.NoError(t, err)

	output := sb.String()
	require.Contains(t, output, "<span")
	require.Contains(t, output, "id=\"id\"")
	require.Contains(t, output, "class=\"class\"")
	require.Contains(t, output, "text")
	require.Contains(t, output, "</span>")
}
