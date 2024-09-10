package input

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	inchesToMeters = 0.0254
	poundsToKg     = 0.453592
)

// ParseInput parses weight and height, converts based on the units.
func ParseInput(weightStr, heightStr, weightUnit, heightUnit string) (float64, float64, error) {
	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid weight value")
	}

	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid height value")
	}

	// Convert weight to kilograms
	if strings.ToLower(weightUnit) == "lbs" {
		weight *= poundsToKg
	}

	// Convert height to meters
	if strings.ToLower(heightUnit) == "in" {
		height *= inchesToMeters
	} else {
		height /= 100 // convert cm to meters
	}

	return weight, height, nil
}
