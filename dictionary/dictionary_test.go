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
	t.Run("update the value for same key", func(t *testing.T) {
		dictionary := Dictionary{"hello": "world"}
		want := "key updated"
		got, err := dictionary.Update("hello", "Peter")
		if err != nil {
			t.Fatal("word not updated man", err)
		}
		AssertStrings(t, got, want)

		definition, err := dictionary.Search("hello")
		if err != nil {
			t.Fatal("could not find updated word:", err)
		}
		AssertStrings(t, definition, "Peter")
	})
	t.Run("update the value but key not there", func(t *testing.T) {
		dictionary := Dictionary{"hello": "world"}
		want := "key not found"
		_, err := dictionary.Update("Parker", "Peter")
		if err == nil {
			t.Fatal("expected an error but got nil", err)
		}
		AssertStrings(t, err.Error(), want)
	})
	t.Run("deleting the existing key-value pair ", func(t *testing.T) {
		dictionary := Dictionary{"hello": "world"}
		want := "key deleted"
		got, err := dictionary.Delete("hello")
		if err != nil {
			t.Fatal("could not find word:", err)
		}
		AssertStrings(t,got,want)
	})
	t.Run("deleting the word not present", func(t *testing.T) {
		dictionary := Dictionary{}
		want:="key not found"
		got,err:=dictionary.Delete("hello")
		if err==nil{
			t.Fatal("expected error but got nil")
		}
		AssertStrings(t,got,want)

	})

}
func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}

}
