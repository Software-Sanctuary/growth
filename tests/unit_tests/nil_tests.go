package unit_tests

import (
	"testing"
)

func Nil(t *testing.T, obj interface{}) {
	if obj != nil {
		t.Error()
	}
}

func NotNil(t *testing.T, obj interface{}) {
	if obj == nil {
		t.Error()
	}
}
