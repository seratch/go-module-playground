package hello

import "testing"

func TestVersion(t *testing.T) {
	want := "0.1.0"
	if got := Version(); got != want {
		t.Errorf("Version() = %q, want %q", got, want)
	}
}
