package sydercomponents_test

import (
	"strings"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	sydercomponents "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	h "maragu.dev/gomponents/html"
)

func TestComponent_SimpleOptions(t *testing.T) {
	testcases := []struct {
		name  string
		comp  sc.Component
		check func(t *testing.T, comp sc.Component)
	}{
		{
			name: "test Class",
			comp: sc.Div(sc.Class("flex")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "flex", comp.GetClass())
			},
		},
		{
			name: "test Class overwrites",
			comp: sc.Div(sc.Class("flex"), sc.Class("none")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "none", comp.GetClass())
				require.NotEqual(t, "flex", comp.GetClass())
			},
		},
		{
			name: "test Id",
			comp: sc.Div(sc.ID("id")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "id", comp.GetID())
			},
		},
		{
			name: "test Data",
			comp: sc.Div(sc.Data("apple", "1"), sc.Data("token", "secret")),
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
			name: "test Aria",
			comp: sc.Div(sc.Aria("apple", "1"), sc.Aria("token", "secret")),
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
		{
			name: "test Kind",
			comp: sc.Div(),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, sc.KindElement, comp.GetKind())
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
			comp := sc.Div(testcase.optFn(testcase.attrValue))
			v, has := testcase.getter(comp)
			require.Equal(t, testcase.attrValue, v)
			require.True(t, has)

			comp = sc.Div()
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
			name: "test TabIndex",
			comp: sc.Div(sc.TabIndex(-1)),
			check: func(t *testing.T, comp sc.Component) {
				v, has := comp.GetTabIndex()
				require.Equal(t, -1, v)
				require.True(t, has)
			},
		},
		{
			name: "test TabIndex missing",
			comp: sc.Div(),
			check: func(t *testing.T, comp sc.Component) {
				v, has := comp.GetTabIndex()
				require.Empty(t, v)
				require.False(t, has)
			},
		},
		{
			name: "test ClassAppend",
			comp: sc.Div(sc.ClassAppend("flex"), sc.ClassAppend("flex-col")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "flex flex-col", comp.GetClass())
			},
		},
		{
			name: "test Class and ClassAppend",
			comp: sc.Div(sc.Class("flex"), sc.ClassAppend("flex-col")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "flex flex-col", comp.GetClass())
			},
		},
		{
			name: "test ClassAppend on empty class",
			comp: sc.Div(sc.ClassAppend("  flex "), sc.ClassAppend(" flex-col")),
			check: func(t *testing.T, comp sc.Component) {
				require.Equal(t, "flex flex-col", comp.GetClass())
			},
		},
		{
			name: "test Disabled",
			comp: sc.Div(sc.Disabled()),
			check: func(t *testing.T, comp sc.Component) {
				require.True(t, comp.HasDisabled())
			},
		},
		{
			name: "test Disabled missing",
			comp: sc.Div(),
			check: func(t *testing.T, comp sc.Component) {
				require.False(t, comp.HasDisabled())
			},
		},
		{
			name: "test Required",
			comp: sc.Div(sc.Required()),
			check: func(t *testing.T, comp sc.Component) {
				require.True(t, comp.HasRequired())
			},
		},
		{
			name: "test Required missing",
			comp: sc.Div(),
			check: func(t *testing.T, comp sc.Component) {
				require.False(t, comp.HasRequired())
			},
		},
		{
			name: "test Hidden",
			comp: sc.Div(sc.Hidden()),
			check: func(t *testing.T, comp sc.Component) {
				require.True(t, comp.HasHidden())
			},
		},
		{
			name: "test Hidden missing",
			comp: sc.Div(),
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
		component := sc.Div(
			sc.Children(
				sc.Div(),
				sc.Div(
					sc.Children(
						sc.Div(),
						sc.Div(),
					),
				),
				sc.Div(),
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

	t.Run("Children overwrite", func(t *testing.T) {
		component := sc.Div(
			sc.Children(sc.Div(), sc.Div()),
			sc.Children(sc.Div()),
		)

		require.Len(t, component.GetChildren(), 1, "should have 1 child on overwrite")
	})

	t.Run("Children and ChildrenAppend", func(t *testing.T) {
		component := sc.Div(
			sc.Children(sc.Div(), sc.Div()),
			sc.ChildrenAppend(sc.Div()),
		)

		require.Len(t, component.GetChildren(), 3, "should have 3 children with append")
	})

	t.Run("Children with Nodes interleaving", func(t *testing.T) {
		component := sc.Div(
			sc.Children(
				sc.Div(sc.ID("child1")),
				sc.Node(h.Div(h.ID("child2"))),
				sc.Div(sc.ID("child3")),
			),
		)

		require.Len(t, component.GetChildren(), 3)

		// 1st child is Element Component
		require.Equal(t, "child1", component.GetChildren()[0].GetID())
		require.Equal(t, sc.KindElement, component.GetChildren()[0].GetKind())

		// 2nd child is Node Component
		require.Equal(t, sc.KindNode, component.GetChildren()[1].GetKind())

		// 3rd child is Element Component
		require.Equal(t, "child3", component.GetChildren()[2].GetID())
		require.Equal(t, sc.KindElement, component.GetChildren()[2].GetKind())
	})
}

func TestComponent_Render(t *testing.T) {
	comp := sc.Span(
		sc.ID("id"),
		sc.Class("class"),
		sc.Children(sc.Text("text")),
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
