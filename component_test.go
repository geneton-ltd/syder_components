package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
	// g "maragu.dev/gomponents"
)

func TestComponent_Children(t *testing.T) {
	component := sc.Any("div",
		sc.WithChildren(
			sc.Any("div"),
			sc.Any("div",
				sc.WithChildren(
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
}
