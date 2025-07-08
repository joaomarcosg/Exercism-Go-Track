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

func TestHappyBirthday(t *testing.T) {
	expected := "Happy birthday Frank! You are now 58 years old!"
	got := HappyBirthday("Frank", 58)

	if got != expected {
		t.Errorf("HappyBirthday() = %s, want %s", got, expected)
	}
}
