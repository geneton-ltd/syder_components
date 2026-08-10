package sydercomponents_test

import (
	"strings"
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestCard_Options(t *testing.T) {
	// create card component with options
	comp := sc.Card(sc.WithID("id"), sc.WithData("name", "john"))

	// test options passthrough with getters
	require.Equal(t, "id", comp.GetID())
	v, has := comp.GetData("name")
	require.Equal(t, "john", v)
	require.True(t, has)
}

func hasCardClasses(t *testing.T, class string) {
	require.Contains(t, class, "w-full")
	require.Contains(t, class, "rounded-md")
	require.Contains(t, class, "shadow-xs")
	require.Contains(t, class, "p-2")
	require.Contains(t, class, "border")
	require.Contains(t, class, "border-1")
	require.Contains(t, class, "border-gray-600")
}

func TestCard_DefaultStyle(t *testing.T) {
	comp := sc.Card()
	hasCardClasses(t, comp.GetClass())
}

func TestCard_Output(t *testing.T) {
	comp := sc.Card()

	var sb strings.Builder
	err := comp.Render(&sb)
	require.NoError(t, err)
	output := sb.String()

	// check if its a <div>
	require.Contains(t, output, "<div")
}

func TestSubCard_Merging(t *testing.T) {
	comp := sc.SubCard(sc.WithClassAppend("custom-class"))
	class := comp.GetClass()

	// has the default Card style
	hasCardClasses(t, class)

	// has its own style
	require.Contains(t, class, "border-dashed")
	require.Contains(t, class, "hover:bg-white")

	// has custom class
	require.Contains(t, class, "custom-class")
}
