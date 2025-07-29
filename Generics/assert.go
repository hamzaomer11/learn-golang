package generics

import "testing"

func AssertEqual(t *testing.T, got, expected interface{}) {
	t.Helper()
	if got != expected {
		t.Errorf("got %+v, expected %+v", got, expected)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, expected T) {
	t.Helper()
	if got == expected {
		t.Errorf("got %+v, expected %+v", got, expected)
	}
}

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
