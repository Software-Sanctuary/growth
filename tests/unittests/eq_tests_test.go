package unittests

import (
	"testing"
)

type FirstObject struct {
	A string
	B int
	C float32
	D InnerObject
}

type SecondObject struct {
	A string
	B int
	C float32
	D InnerObject
}

type InnerObject struct {
	Yipy string
}

func InitEqualObjects() (FirstObject, FirstObject) {
	return FirstObject{
			A: "Test",
			B: 1,
			C: 1.0,
			D: InnerObject{
				Yipy: "Inner",
			},
		},
		FirstObject{
			A: "Test",
			B: 1,
			C: 1.0,
			D: InnerObject{
				Yipy: "Inner",
			},
		}
}

func InitNonEqualObjects() (FirstObject, SecondObject) {
	return FirstObject{
			A: "Test",
			B: 1,
			C: 1.0,
			D: InnerObject{
				Yipy: "Inner",
			},
		},
		SecondObject{
			A: "Test",
			B: 1,
			C: 1.0,
			D: InnerObject{
				Yipy: "Inner",
			},
		}
}

func TestEqual(t *testing.T) {
	frst, sec := InitEqualObjects()
	t.Run(
		" objects should equal",
		func(t *testing.T) {
			Equal(t, frst, sec)
		},
	)
}

func TestNotEqual(t *testing.T) {
	frst, sec := InitNonEqualObjects()
	t.Run(
		" objects should equal",
		func(t *testing.T) {
			NotEqual(t, frst, sec)
		},
	)
}
