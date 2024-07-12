package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello to people", func(t *testing.T) {
		got := hello("Chris", "English")
		want := "Hello, Chris"
		assert(t, got, want)
	})

	t.Run("hello when empty string", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, world"

		assert(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assert(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := hello("dannis", "French")
		want := "Bonjur, dannis"

		assert(t, got, want)
	})
}

func assert(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
