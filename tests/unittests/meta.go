package unittests

import "testing"

type UnitTest struct {
	Description string
	Function    func(t *testing.T)
}

type UnitTests []UnitTest

type UnitTesting struct {
	Tests UnitTests
}
