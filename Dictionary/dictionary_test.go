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
	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)

	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)

	})

}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is a new test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)

}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
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

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	assertStrings(t, got, definition)

}
