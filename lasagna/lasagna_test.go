package lasagna

import "testing"

func TestRemaningOvenTime(t *testing.T) {
	expected := 30
	got := RemaningOvenTime(10)

	if expected != got {
		t.Errorf("RemaningOvenTime() = %q; want %q", got, expected)
	}
}

func TestPrepationTime(t *testing.T) {
	expected := 4
	got := PreparationTime(2)

	if expected != got {
		t.Errorf("PreparationTime() = %q; want %q", got, expected)
	}
}

func TestElepsedTime(t *testing.T) {
	expected := 26
	got := ElepsedTime(3, 20)

	if expected != got {
		t.Errorf("ElepsedTime() = %q; want %q", got, expected)
	}
}
