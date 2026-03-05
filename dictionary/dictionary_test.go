package dictionary

import (
	"testing"
)

func TestSeach(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	t.Run("add new word", func(t *testing.T) {
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("add existing word", func(t *testing.T) {

		err := dictionary.Add(word, "new test definition")
		assertError(t, err, ErrWordExists)

		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	t.Run("update existing word", func(t *testing.T) {
		updatedDefinition := "updated definition"
		dictionary.Update(word, updatedDefinition)

		assertDefinition(t, dictionary, word, updatedDefinition)
	})
	t.Run("update word doesn't exist", func(t *testing.T) {
		nonEntryWord := "uhoh"
		err := dictionary.Update(nonEntryWord, "uh-oh")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "delete-test"
	definition := "this is just a test for deletion"
	dictionary := Dictionary{word: definition}
	t.Run("delete word", func(t *testing.T) {
		err := dictionary.Delete(word)

		assertError(t, err, nil)

		_, findErr := dictionary.Search(word)

		assertError(t, findErr, ErrNotFound)

	})
	t.Run("delete word not exists", func(t *testing.T) {
		err := dictionary.Delete("uhoh")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, got, definition)
}
