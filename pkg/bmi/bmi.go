package bmi

// CalculateBMI calculates the BMI using weight and height
func CalculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

// InterpretBMI gives a description based on the BMI value
func InterpretBMI(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi <= 24.9:
		return "Normal weight"
	case bmi <= 29.9:
		return "Overweight"
	default:
		return "Obese"
	}
}
