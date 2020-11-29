package data

import "testing"

func TestGet(t *testing.T) {
	want := "monsters"
	if got := Get(want); got != want {
		t.Errorf("Get(%q) = %q, want %q", want, got, want)
	}
}
