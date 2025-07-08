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

func TestAssignTable(t *testing.T) {
	expected := "Welcome to my party Christiane!You have been assigned to table 027.Your table is on the left, exactly 23.8 meters from here. You will be sitting next to Frank"
	got := AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298)

	if got != expected {
		t.Errorf("AssingTable() = %s, want %s", got, expected)
	}
}
