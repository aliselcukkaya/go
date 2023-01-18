package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		AssertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		AssertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "addtest"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		AssertError(t, err, nil)
		AssertDefinition(t, dictionary, word, definition)
	})
	t.Run("already exists", func(t *testing.T) {
		word := "addtest"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)

		AssertError(t, err, ErrAlreadyExists)
		AssertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "updatetest"
		definition := "will be updated"
		dictionary := Dictionary{word: definition}
		updatedDefinition := "updated definition"

		err := dictionary.Update(word, updatedDefinition)

		AssertError(t, err, nil)
		AssertDefinition(t, dictionary, word, updatedDefinition)
	})
	t.Run("non existing word", func(t *testing.T) {
		word := "non existing word"
		definition := "non existing definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		AssertError(t, err, ErrWordDoesNotExist)
	})
}

func AssertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func AssertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
