package helloworld

import "testing"

func TestGreeting(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Greeting("Adam", "")
		want := "Hello World, Adam"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Greeting("", "")
		want := "Hello World, Kind Stranger"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Greeting("Elie", "Spanish")
		want := "Hola Amigo, Elie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
