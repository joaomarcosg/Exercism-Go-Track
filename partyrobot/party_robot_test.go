package partyrobot

import (
	"testing"
)

func TestWelcome(t *testing.T) {
	expected := "Welcome to my party Christiane"
	got := Welcome("Christiane")

	if got != expected {
		t.Errorf("Welcome() = %s, want %s", got, expected)
	}
}
