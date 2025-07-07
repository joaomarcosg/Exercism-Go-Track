package partyrobot

import "testing"

func TestWelcome(t *testing.T) {
	expected := "Christiane"
	got := Welcome("Christiane")

	if got != expected {
		t.Errorf("Welcome() = %q; want %q", got, expected)
	}
}
