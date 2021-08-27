package mapper

import (
	"testing"

	"github.com/gokhantamkoc/growth-framework/tests/unittests"
)

type A struct {
	Test string `test:"A" dangle:"B"`
}

type B struct {
	Test string
}

func TestReadTag(t *testing.T) {
	res, err := ReadTag(&A{}, "Test", "test")

	t.Run("err should be null", func(t *testing.T) {
		unittests.Nil(t, err)
	})

	t.Run("result should be equal to A", func(t *testing.T) {
		unittests.Equal(t, "A", res)
	})
}

func TestMap(t *testing.T) {
	a := A{Test: "Test"}
	b := B{}
	MapObject(a, b)
	t.Run("result should be equal", func(t *testing.T) {
		unittests.Equal(t, a.Test, b.Test)
	})
}
