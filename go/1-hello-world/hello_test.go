package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q wamt %q", got, want)
		}
	}

	t.Run("say hello", func(t *testing.T) {
		got := Hello("world!", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("default to hello world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in indonesia", func(t *testing.T) {
		got := Hello("andrkrn!", "id")
		want := "Hai, andrkrn!"

		assertCorrectMessage(t, got, want)
	})
}
