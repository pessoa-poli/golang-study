package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people.", func(t *testing.T) {
		have := Hello("Jack", "")
		want := "Hello, Jack"
		assertCorrectMessage(t, have, want)
	})
	t.Run("Testing call to empty hello", func(t *testing.T) {
		have := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, have, want)
	})
	t.Run("passing a language", func(t *testing.T) {
		have := Hello("Jake", "Spanish")
		want := "Hola, Jake"
		assertCorrectMessage(t, have, want)
	})
	t.Run("Greeting in french", func(t *testing.T) {
		have := Hello("Jake", "French")
		want := "Bonjour, Jake"
		assertCorrectMessage(t, have, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
