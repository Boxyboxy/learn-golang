package generics

import "testing"

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

// to write generic functions in go, you need to provide type parameters which is just a fancy way of saying
// decsribe your generic type and give it a label
// type of our type parameter is comparable
// we're using comparable because we want to describe to the compilaer that we wish to sue the == and != operators
// on things of type T
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

// note: There is a type of type parameter called any
// any is just an alias for interface{}
// func GenericFoo[T any](x, y T)             similar to interface{} but will only work for GenericFoo(apple1, apple2)
// func InterfaceyFoo(x, y interface{})       slightly different from any, will work with InterfaceyFoo(apple, orange)
