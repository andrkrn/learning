package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat correctly", func(t *testing.T) {
		got := Repeat("a", -1)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("correct repeat times", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
