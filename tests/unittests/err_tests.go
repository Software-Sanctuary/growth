package unittests

import "testing"

func Error(t *testing.T, err error) {
	if err == nil {
		t.Error()
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Error()
	}
}
