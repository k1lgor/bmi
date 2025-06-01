package bmi

// CalculateBMI returns the Body Mass Index (BMI) given weight (kg) and height (m).
// Formula: BMI = weight / (height * height)
func CalculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

// InterpretBMI returns a human-readable BMI category string
// based on standard BMI classification ranges.
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
