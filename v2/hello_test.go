package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestVersion(t *testing.T) {
	want := "2.0.0"
	if got := Version(); got != want {
		t.Errorf("Version() = %q, want %q", got, want)
	}
}
