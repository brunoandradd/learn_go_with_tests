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
	dic := Dictionary{}
	dic.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dic.Search("test")
	if err != nil {
		t.Fatal("should find add word:", err)
	}

	assertStrings(t, got, want)
}

func TestToString(t *testing.T) {
	d := Dollar(20)
	got := d.ToSring()
	want := "$ 20"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// continue here
// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps#pointers-copies-et-al
