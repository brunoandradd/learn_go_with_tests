package maps

import "testing"

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dic.Search("wnknown")
		want := "word not found"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		word := "test"
		want := "this is just a test"

		dic.Add(word, "this is just a test")

		assertDefinition(t, dic, word, want)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		want := "this is just a test"
		dic := Dictionary{word: want}

		err := dic.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, want)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		want := "this is just a test"
		dic := Dictionary{word: want}
		newWant := "update value"
		err := dic.Update(word, newWant)

		assertError(t, err, nil)
		assertDefinition(t, dic, word, newWant)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		want := "this is just a test"
		dic := Dictionary{}
		err := dic.Update(word, want)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		word := "test"
		dic := Dictionary{word: "test definition"}

		dic.Delete(word)

		_, err := dic.Search("test")

		assertError(t, err, ErrNotFound)
	})
}

func TestToString(t *testing.T) {
	d := Dollar(20)
	got := d.ToSring()
	want := "$ 20"

	assertStrings(t, got, want)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("expect error type %v", want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
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
