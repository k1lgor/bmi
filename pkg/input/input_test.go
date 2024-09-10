package input

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	weight, height, err := ParseInput("70", "170", "kg", "cm")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if weight != 70 || height != 1.7 {
		t.Errorf("Expected weight 70kg and height 1.7m, got weight %.2fkg and height %.2fm", weight, height)
	}

	// Test invalid input
	_, _, err = ParseInput("invalid", "170", "kg", "cm")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
