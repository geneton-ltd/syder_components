package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

// TableContainer
func TestTableContainer_Options(t *testing.T) {
	comp := sc.TableContainer(sc.WithID("id"), sc.WithData("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasTableContainerClasses(t *testing.T, class string) {
	require.Contains(t, class, "overflow-x-auto")
	require.Contains(t, class, "rounded")
	require.Contains(t, class, "border")
	require.Contains(t, class, "border-gray-300")
	require.Contains(t, class, "shadow-sm")
}

func TestTableContainer_DefaultStyle(t *testing.T) {
	comp := sc.TableContainer()
	hasTableContainerClasses(t, comp.GetClass())
}

func TestTableContainer_Output(t *testing.T) {
	id := "id"
	comp := sc.TableContainer(sc.WithID(id))
	assertElementWithIDExists(t, comp, "div", id)
}

// Table
func TestTable_Options(t *testing.T) {
	comp := sc.Table(sc.WithID("id"), sc.WithData("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasTableClasses(t *testing.T, class string) {
	require.Contains(t, class, "min-w-full")
	require.Contains(t, class, "divide-y-2")
	require.Contains(t, class, "divide-gray-200")
}

func TestTable_DefaultStyle(t *testing.T) {
	comp := sc.Table()
	hasTableClasses(t, comp.GetClass())
}

func TestTable_Output(t *testing.T) {
	id := "id"
	comp := sc.Table(sc.WithID(id))
	assertElementWithIDExists(t, comp, "table", id)
}

// THead
func TestTHead_Options(t *testing.T) {
	comp := sc.THead(sc.WithID("id"), sc.WithData("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func TestTHead_DefaultStyle(t *testing.T) {
	comp := sc.THead()
	// THead currently has no default classes
	require.Empty(t, comp.GetClass())
}

func TestTHead_Output(t *testing.T) {
	id := "id"
	comp := sc.THead(sc.WithID(id))
	wrapper := sc.Table(sc.WithChildren(comp))
	assertElementWithIDExists(t, wrapper, "thead", id)
}

// TBody
func TestTBody_Options(t *testing.T) {
	comp := sc.TBody(sc.WithID("id"), sc.WithData("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasTBodyClasses(t *testing.T, class string) {
	require.Contains(t, class, "divide-y")
	require.Contains(t, class, "divide-gray-200")
}

func TestTBody_DefaultStyle(t *testing.T) {
	comp := sc.TBody()
	hasTBodyClasses(t, comp.GetClass())
}

func TestTBody_Output(t *testing.T) {
	id := "id"
	comp := sc.TBody(sc.WithID(id))
	wrapper := sc.Table(sc.WithChildren(comp))
	assertElementWithIDExists(t, wrapper, "tbody", id)
}

// Tr
func TestTr_Options(t *testing.T) {
	comp := sc.Tr(sc.WithID("id"), sc.WithData("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasTrClasses(t *testing.T, class string) {
	require.Contains(t, class, "hover:bg-gray-100")
}

func TestTr_DefaultStyle(t *testing.T) {
	comp := sc.Tr()
	hasTrClasses(t, comp.GetClass())
}

func TestTr_Output(t *testing.T) {
	id := "id"
	comp := sc.Tr(sc.WithID(id))
	wrapper := sc.Table(sc.WithChildren(comp))
	assertElementWithIDExists(t, wrapper, "tr", id)
}

// Th
func TestTh_Options(t *testing.T) {
	comp := sc.Th(sc.WithID("id"), sc.WithData("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasThClasses(t *testing.T, class string) {
	require.Contains(t, class, "px-3")
	require.Contains(t, class, "py-2")
	require.Contains(t, class, "whitespace-nowrap")
	require.Contains(t, class, "font-medium")
	require.Contains(t, class, "text-gray-900")
}

func TestTh_DefaultStyle(t *testing.T) {
	comp := sc.Th()
	hasThClasses(t, comp.GetClass())
}

func TestTh_Output(t *testing.T) {
	id := "id"
	comp := sc.Th(sc.WithID(id))
	wrapper := sc.Table(sc.WithChildren(comp))
	assertElementWithIDExists(t, wrapper, "th", id)
}

// Td
func TestTd_Options(t *testing.T) {
	comp := sc.Td(sc.WithID("id"), sc.WithData("name", "john"))

	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasTdClasses(t *testing.T, class string) {
	require.Contains(t, class, "px-3")
	require.Contains(t, class, "py-2")
	require.Contains(t, class, "whitespace-nowrap")
	require.Contains(t, class, "text-gray-900")
}

func TestTd_DefaultStyle(t *testing.T) {
	comp := sc.Td()
	hasTdClasses(t, comp.GetClass())
}

func TestTd_Output(t *testing.T) {
	id := "id"
	comp := sc.Td(sc.WithID(id))
	wrapper := sc.Table(sc.WithChildren(comp))
	assertElementWithIDExists(t, wrapper, "td", id)
}
