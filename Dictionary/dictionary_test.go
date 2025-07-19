package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, got, expected)

	})

	t.Run("unknown word", func(t *testing.T) {

		_, got := dictionary.Search("unknown")
		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, got, ErrNotFound)

	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	expected := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, expected)
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
