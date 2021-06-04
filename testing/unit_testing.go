package testing

import (
	"testing"
)

type UnitTest struct {
	Description string
	Function    func(t *testing.T)
}

type UnitTests []UnitTest

type UnitTesting struct {
	Tests UnitTests
}

// func (unitTesting UnitTesting) test(t *testing.T) {
//	for index, ut := range unitTesting.Tests {
//		t.Run(fmt.Sprintf("%d. %s", index+1, ut.Description), ut.Function)
//	}
// }
