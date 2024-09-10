package bmi

import "testing"

func TestCalculateBMI(t *testing.T) {
	bmi := CalculateBMI(70, 1.75)
	expected := 22.86
	tolerance := 0.01

	if (bmi-expected) > tolerance || (expected-bmi) > tolerance {
		t.Errorf("Expected %.2f, got %.2f", expected, bmi)
	}
}

func TestInterpretBMI(t *testing.T) {
	tests := []struct {
		bmi      float64
		expected string
	}{
		{17.5, "Underweight"},
		{22.0, "Normal weight"},
		{27.0, "Overweight"},
		{32.0, "Obese"},
	}

	for _, test := range tests {
		result := InterpretBMI(test.bmi)
		if result != test.expected {
			t.Errorf("For BMI %.2f, expected %s, got %s", test.bmi, test.expected, result)
		}
	}
}
