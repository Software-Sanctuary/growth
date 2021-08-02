package mapper

import (
	"github.com/gokhantamkoc/growth-framework/tests/unit_tests"
	"testing"
)

type A struct {
	Test string `test:"A" dangle:"B"`
}

func TestReadTag(t *testing.T) {
	res, err := ReadTag(&A{}, "Test", "test")
	unit_tests.NotNil(t, err)
	unit_tests.Equal(t, res, "A")
}
