package hello

import "testing"

// TestHello tests Hello()
func TestHello(t *testing.T) {
	got := Hello()
	want := "micro"
	if got != want {
		t.Errorf("got %q and want %q", got, want)
	}
}
