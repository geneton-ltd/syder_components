package components_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestTableComponents(t *testing.T) {
	testcases := []struct {
		name            string
		tag             string
		comp            func(children ...sc.Node) sc.Node
		cls             string
		shouldBeInTable bool
	}{
		{
			name: "table container",
			tag:  "div",
			comp: co.TableContainer,
			cls:  "overflow-x-auto rounded border border-gray-300 shadow-sm",
		},
		{
			name: "table",
			tag:  "table",
			comp: co.Table,
			cls:  "min-w-full divide-y-2 divide-gray-200",
		},
		{
			name:            "thead",
			tag:             "thead",
			comp:            co.THead,
			cls:             "",
			shouldBeInTable: true,
		},
		{
			name:            "tbody",
			tag:             "tbody",
			comp:            co.TBody,
			cls:             "divide-y divide-gray-200",
			shouldBeInTable: true,
		},
		{
			name:            "tr",
			tag:             "tr",
			comp:            co.Tr,
			cls:             "hover:bg-gray-100",
			shouldBeInTable: true,
		},
		{
			name:            "th",
			tag:             "th",
			comp:            co.Th,
			cls:             "px-3 py-2 whitespace-nowrap font-medium text-gray-900",
			shouldBeInTable: true,
		},
		{
			name:            "td",
			tag:             "td",
			comp:            co.Td,
			cls:             "px-3 py-2 whitespace-nowrap text-gray-900",
			shouldBeInTable: true,
		},
	}

	for _, tt := range testcases {
		// test options passthrough with getter
		t.Run(tt.name+" options", func(t *testing.T) {
			comp := tt.comp(sc.ID("id"), sc.Data("name", "john"))

			v, has := comp.Attr("id")
			require.True(t, has)
			require.Equal(t, "id", v)

			v, has = comp.Attr("data-name")
			require.True(t, has)
			require.Equal(t, "john", v)
		})

		t.Run(tt.name+" default style", func(t *testing.T) {
			comp := tt.comp()
			cls, hasCls := comp.Attr("class")
			if tt.cls != "" {
				require.True(t, hasCls)
			} else {
				require.False(t, hasCls)
			}
			require.Equal(t, tt.cls, cls)
		})

		t.Run(tt.name+" output", func(t *testing.T) {
			id := "id"
			comp := tt.comp(sc.ID(id))

			if tt.shouldBeInTable {
				comp = co.Table(comp)
			}

			tu.AssertElementWithIDExists(t, comp, tt.tag, id)
		})
	}
}
