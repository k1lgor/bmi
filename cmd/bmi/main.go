package main

import (
	"bmi/pkg/bmi"
	"bmi/pkg/input"
	"bmi/ui"
	"flag"
	"fmt"
	"log"
	"os"
)

const version = "0.4.0"

func main() {
	// Flag parsing
	showVersion := flag.Bool("version", false, "Display the version")
	useTUI := flag.Bool("tui", false, "Run in TUI mode")
	flag.Parse()

	// Show version and exit if --version is provided
	if *showVersion {
		fmt.Println("BMI Calculator version:", version)
		os.Exit(0)
	}

	// Run TUI if --tui is provided
	if *useTUI {
		err := ui.StartTUI()
		if err != nil {
			log.Fatalf("Error running TUI: %v", err)
		}
		return
	}

	// CLI mode expects 4 arguments: <weight> <kg|lbs> <height> <cm|in>
	if len(os.Args) != 5 {
		log.Fatal("Usage: bmi <weight> <kg|lbs> <height> <cm|in> or use --tui for interactive mode.")
	}

	weightStr, weightUnit, heightStr, heightUnit := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	// Parse inputs
	weight, height, err := input.ParseInput(weightStr, heightStr, weightUnit, heightUnit)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// Calculate BMI
	bmiValue := bmi.CalculateBMI(weight, height)
	result := bmi.InterpretBMI(bmiValue)

	// Output results
	fmt.Printf("Your BMI: %.2f\nCategory: %s\n", bmiValue, result)
}
