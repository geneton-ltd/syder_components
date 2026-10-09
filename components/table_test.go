package components_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

// TableContainer
func TestTableContainer_Options(t *testing.T) {
	comp := co.TableContainer(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func hasTableContainerClasses(t *testing.T, class string) {
	require.Contains(t, class, "overflow-x-auto")
	require.Contains(t, class, "rounded")
	require.Contains(t, class, "border")
	require.Contains(t, class, "border-gray-300")
	require.Contains(t, class, "shadow-sm")
}

func TestTableContainer_DefaultStyle(t *testing.T) {
	comp := co.TableContainer()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasTableContainerClasses(t, class)
}

func TestTableContainer_Output(t *testing.T) {
	id := "id"
	comp := co.TableContainer(sc.ID(id))
	tu.AssertElementWithIDExists(t, comp, "div", id)
}

// Table
func TestTable_Options(t *testing.T) {
	comp := co.Table(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func hasTableClasses(t *testing.T, class string) {
	require.Contains(t, class, "min-w-full")
	require.Contains(t, class, "divide-y-2")
	require.Contains(t, class, "divide-gray-200")
}

func TestTable_DefaultStyle(t *testing.T) {
	comp := co.Table()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasTableClasses(t, class)
}

func TestTable_Output(t *testing.T) {
	id := "id"
	comp := co.Table(sc.ID(id))
	tu.AssertElementWithIDExists(t, comp, "table", id)
}

// THead
func TestTHead_Options(t *testing.T) {
	comp := co.THead(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func TestTHead_DefaultStyle(t *testing.T) {
	comp := co.THead()
	class, hasClass := comp.Attr("class")
	require.False(t, hasClass)
	// THead currently has no default classes
	require.Empty(t, class)
}

func TestTHead_Output(t *testing.T) {
	id := "id"
	comp := co.THead(sc.ID(id))
	wrapper := co.Table(comp)
	tu.AssertElementWithIDExists(t, wrapper, "thead", id)
}

// TBody
func TestTBody_Options(t *testing.T) {
	comp := co.TBody(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func hasTBodyClasses(t *testing.T, class string) {
	require.Contains(t, class, "divide-y")
	require.Contains(t, class, "divide-gray-200")
}

func TestTBody_DefaultStyle(t *testing.T) {
	comp := co.TBody()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasTBodyClasses(t, class)
}

func TestTBody_Output(t *testing.T) {
	id := "id"
	comp := co.TBody(sc.ID(id))
	wrapper := co.Table(comp)
	tu.AssertElementWithIDExists(t, wrapper, "tbody", id)
}

// Tr
func TestTr_Options(t *testing.T) {
	comp := co.Tr(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func hasTrClasses(t *testing.T, class string) {
	require.Contains(t, class, "hover:bg-gray-100")
}

func TestTr_DefaultStyle(t *testing.T) {
	comp := co.Tr()

	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasTrClasses(t, class)
}

func TestTr_Output(t *testing.T) {
	id := "id"
	comp := co.Tr(sc.ID(id))
	wrapper := co.Table(comp)
	tu.AssertElementWithIDExists(t, wrapper, "tr", id)
}

// Th
func TestTh_Options(t *testing.T) {
	comp := co.Th(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func hasThClasses(t *testing.T, class string) {
	require.Contains(t, class, "px-3")
	require.Contains(t, class, "py-2")
	require.Contains(t, class, "whitespace-nowrap")
	require.Contains(t, class, "font-medium")
	require.Contains(t, class, "text-gray-900")
}

func TestTh_DefaultStyle(t *testing.T) {
	comp := co.Th()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasThClasses(t, class)
}

func TestTh_Output(t *testing.T) {
	id := "id"
	comp := co.Th(sc.ID(id))
	wrapper := co.Table(comp)
	tu.AssertElementWithIDExists(t, wrapper, "th", id)
}

// Td
func TestTd_Options(t *testing.T) {
	comp := co.Td(sc.ID("id"), sc.Data("name", "john"))

	// test options passthrough with getters
	v, has := comp.Attr("id")
	require.True(t, has)
	require.Equal(t, "id", v)

	v, has = comp.Attr("data-name")
	require.True(t, has)
	require.Equal(t, "john", v)
}

func hasTdClasses(t *testing.T, class string) {
	require.Contains(t, class, "px-3")
	require.Contains(t, class, "py-2")
	require.Contains(t, class, "whitespace-nowrap")
	require.Contains(t, class, "text-gray-900")
}

func TestTd_DefaultStyle(t *testing.T) {
	comp := co.Td()
	class, hasClass := comp.Attr("class")
	require.True(t, hasClass)
	hasTdClasses(t, class)
}

func TestTd_Output(t *testing.T) {
	id := "id"
	comp := co.Td(sc.ID(id))
	wrapper := co.Table(comp)
	tu.AssertElementWithIDExists(t, wrapper, "td", id)
}
