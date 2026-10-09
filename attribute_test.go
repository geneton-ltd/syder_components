package sydercomponents_test

import (
	"bytes"
	"fmt"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestAttribute(t *testing.T) {
	type testcase struct {
		name      string
		attrName  string
		attrValue string
		content   string
		boolean   bool
		comp      func() sc.Node
	}

	simpleAttributes := []struct {
		name string
		comp func(value string) sc.Node
	}{
		{name: "class", comp: sc.Class},
		{name: "id", comp: sc.ID},
		{name: "style", comp: sc.Style},
		{name: "name", comp: sc.Name},
		{name: "type", comp: sc.Type},
		{name: "placeholder", comp: sc.Placeholder},
		{name: "value", comp: sc.Value},
		{name: "src", comp: sc.Src},
		{name: "alt", comp: sc.Alt},
		{name: "href", comp: sc.Href},
		{name: "target", comp: sc.Target},
		{name: "lang", comp: sc.Lang},
		{name: "charset", comp: sc.Charset},
		{name: "rel", comp: sc.Rel},
		{name: "content", comp: sc.Content},
		{name: "popovertarget", comp: sc.PopoverTrigger},
	}
	simpleAttributeTestcases := []testcase{}
	for _, attr := range simpleAttributes {
		simpleAttributeTestcases = append(simpleAttributeTestcases, testcase{
			name:      attr.name + " attribute",
			attrName:  attr.name,
			attrValue: "placeholder",
			comp:      func() sc.Node { return attr.comp("placeholder") },
		})
	}

	booleanAttributes := []struct {
		name string
		comp func() sc.Node
	}{
		{name: "hidden", comp: sc.Hidden},
		{name: "disabled", comp: sc.Disabled},
		{name: "required", comp: sc.Required},
	}
	booleanAttributeTestcases := []testcase{}
	for _, attr := range booleanAttributes {
		booleanAttributeTestcases = append(booleanAttributeTestcases, testcase{
			name:     attr.name + " attribute",
			attrName: attr.name,
			boolean:  true,
			comp:     attr.comp,
		})
	}

	testcases := []testcase{
		{
			name:      "render attribute with value",
			attrName:  "id",
			attrValue: "id123",
			comp:      func() sc.Node { return sc.Attr("id", "id123") },
		},
		{
			name:      "render boolean attribute",
			attrName:  "id",
			attrValue: "",
			boolean:   true,
			comp:      func() sc.Node { return sc.Attr("id", "") },
		},
		{
			name:      "tabindex attribute",
			attrName:  "tabindex",
			attrValue: "67",
			comp:      func() sc.Node { return sc.TabIndex(67) },
		},
		{
			name:      "data attribute",
			attrName:  "data-key",
			attrValue: "placeholder",
			comp:      func() sc.Node { return sc.Data("key", "placeholder") },
		},
		{
			name:      "aria attribute",
			attrName:  "aria-key",
			attrValue: "placeholder",
			comp:      func() sc.Node { return sc.Aria("key", "placeholder") },
		},
	}

	testcases = append(testcases, simpleAttributeTestcases...)
	testcases = append(testcases, booleanAttributeTestcases...)

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			attr := tt.comp()

			require.Equal(t, sc.KindAttr, attr.Kind())
			require.Equal(t, tt.attrName, attr.Name())
			require.Equal(t, tt.attrValue, attr.Value())
			require.Empty(t, attr.Children())
			v, has := attr.Attr(tt.attrName)
			require.True(t, has)
			require.Equal(t, tt.attrValue, v)

			var buf bytes.Buffer
			err := attr.Render(&buf)
			require.NoError(t, err)

			content := buf.String()
			if tt.boolean {
				require.Equal(t, tt.attrName, content)
			} else {
				require.Equal(t, fmt.Sprintf("%s=\"%s\"", tt.attrName, tt.attrValue), content)
			}
		})
	}
}

func TestMultiAttributes(t *testing.T) {
	id := "id"

	testcases := []struct {
		name    string
		comp    func(dialogID string) sc.Node
		content []sc.Node
	}{
		{
			name:    "DialogTrigger",
			comp:    sc.DialogTrigger,
			content: []sc.Node{sc.Attr("commandfor", id), sc.Attr("command", "show-modal")},
		},
		{
			name:    "DialogClose",
			comp:    sc.DialogClose,
			content: []sc.Node{sc.Attr("commandfor", id), sc.Attr("command", "close")},
		},
		{
			name:    "PopoverClose",
			comp:    sc.PopoverClose,
			content: []sc.Node{sc.Attr("popovertarget", id), sc.Attr("popovertargetaction", "hide")},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name+" attribute", func(t *testing.T) {
			attr := tt.comp(id)
			require.Equal(t, sc.KindGroup, attr.Kind())
			require.Equal(t, tt.content, attr.Children())
			for _, child := range attr.Children() {
				require.Equal(t, sc.KindAttr, child.Kind())
			}
		})
	}
}
