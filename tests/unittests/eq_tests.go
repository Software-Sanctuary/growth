package unittests

import (
	"testing"
)

func Equal(t *testing.T, frst interface{}, sec interface{}) {
	if frst != sec {
		t.Error()
	}
}

func NotEqual(t *testing.T, frst interface{}, sec interface{}) {
	if frst == sec {
		t.Error()
	}
}
