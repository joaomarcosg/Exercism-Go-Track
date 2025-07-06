package annalyn

import "testing"

type characterState struct {
	archerIsAwake   bool
	prisonerIsAwake bool
	expected        bool
}

func TestCanSignalPrisoner(t *testing.T) {
	tt := characterState{
		archerIsAwake:   false,
		prisonerIsAwake: false,
		expected:        false,
	}

	if got := CanSignalPrisioner(tt.archerIsAwake, tt.prisonerIsAwake); got != tt.expected {
		t.Errorf("CanSignalPrisoner(%v, %v) = %v; want %v", tt.archerIsAwake, tt.prisonerIsAwake, got, tt.expected)
	}
}
