package gateway

import "testing"

func TestGet(t *testing.T) {
	want := "monsters"
	if got := Get(want); got != want {
		t.Errorf("Get(\"monsters\") = %q, want %q", got, want)
	}
}
