package testing

import "testing"

type TestingBlueprint interface {
	test(t testing.T)
}
