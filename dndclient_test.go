package dndclient

import "testing"

func TestGetMonsters(t *testing.T) {
	want := "monsters"
	if got := GetMonsters(); got != want {
		t.Errorf("GetMonsters() = %q, want %q", got, want)
	}
}
