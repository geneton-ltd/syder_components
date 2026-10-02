package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	"github.com/stretchr/testify/require"
)

func TestIf(t *testing.T) {
	// true
	comp := sc.Div(sc.ID("id"))

	cond := sc.If(true, comp)
	assertElementWithIDExists(t, cond, "div", "id")

	// false
	cond = sc.If(false, comp)
	require.Equal(t, sc.Empty(), cond)
}

func TestIff(t *testing.T) {
	// true
	comp := sc.Div(sc.ID("id"))
	called := false
	f := func() sc.Component {
		called = true
		return comp
	}

	cond := sc.Iff(true, f)
	assertElementWithIDExists(t, cond, "div", "id")
	require.True(t, called)

	// false
	called = false
	cond = sc.Iff(false, f)
	require.Equal(t, sc.Empty(), cond)
	require.False(t, called)
}
