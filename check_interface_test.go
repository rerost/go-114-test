package main_test

import "testing"

func TestInterface(t *testing.T) {
	type A interface {
		MethodA() int
		CommonMethod() int
	}

	type B interface {
		MethodB() int
		CommonMethod() int
	}

	type AB interface {
		A
		B
	}
}
