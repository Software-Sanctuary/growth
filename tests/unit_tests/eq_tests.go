package unit_tests

import "testing"

func Equal(t *testing.T, exp interface{}, act interface{}) {
	if exp != act {
		t.Error()
	}
}

func NotEqual(t *testing.T, exp interface{}, act interface{}) {
	if exp == act {
		t.Error()
	}
}
