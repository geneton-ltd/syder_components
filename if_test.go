package sydercomponents_test

import (
	"testing"

	sc "github.com/geneton-ltd/syder_components"
	tu "github.com/geneton-ltd/syder_components/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestIf(t *testing.T) {
	// true
	comp := sc.Div(sc.ID("id"))

	cond := sc.If(true, comp)
	tu.AssertElementWithIDExists(t, cond, "div", "id")

	// false
	cond = sc.If(false, comp)
	require.Nil(t, cond)
}

func TestIff(t *testing.T) {
	// true
	comp := sc.Div(sc.ID("id"))
	called := false
	f := func() sc.Node {
		called = true
		return comp
	}

	cond := sc.Iff(true, f)
	tu.AssertElementWithIDExists(t, cond, "div", "id")
	require.True(t, called)

	// false
	called = false
	cond = sc.Iff(false, f)
	require.Nil(t, cond)
	require.False(t, called)
}
