package input

import (
	"fmt"
	"strconv"
	"strings"
)

// Conversion constants for imperial to metric units.
const (
	inchesToMeters = 0.0254   // 1 inch = 0.0254 meters
	poundsToKg     = 0.453592 // 1 pound = 0.453592 kilograms
)

// ParseInput parses weight and height from strings, converts units if needed,
// and returns metric values (kg, meters).
// Accepts weightUnit as "kg" or "lbs", heightUnit as "cm" or "in".
// Returns error if parsing fails or units are invalid.
func ParseInput(weightStr, heightStr, weightUnit, heightUnit string) (float64, float64, error) {
	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid weight value")
	}

	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid height value")
	}

	// Convert weight to kilograms if needed
	if strings.ToLower(weightUnit) == "lbs" {
		weight *= poundsToKg
	}

	// Convert height to meters if needed
	if strings.ToLower(heightUnit) == "in" {
		height *= inchesToMeters
	} else {
		height /= 100 // convert cm to meters
	}

	return weight, height, nil
}
