package main

import "testing"

func TestMap(t *testing.T) {
	dictionary := Dictionary{"test": "this is just test"}
	t.Run("Word present in dictionary", func(t *testing.T) {
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("unexpected error", err)
		}
		want := "this is just test"
		AssertStrings(t, got, want)
	})
	t.Run("Word not present in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("hello")
		want := "word not found"
		if err == nil {
			t.Fatal("expected to get error")
		}
		AssertStrings(t, err.Error(), want)
	})
	t.Run("Adding a word to dictionary", func(t *testing.T) {
		dictionary.Add("hello", "world")
		want := "world"
		got, err := dictionary.Search("hello")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		AssertStrings(t, got, want)
	})

}
func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}

}
