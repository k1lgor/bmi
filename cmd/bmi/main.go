package main

import (
	"bmi/pkg/bmi"
	"bmi/pkg/input"
	"bmi/ui"
	"bmi/web"
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const version = "0.5.1"

func main() {
	// Set logrus formatting (optional: JSONFormatter for machine logs)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Custom usage/help output
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "BMI Calculator version: %s\n", version)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bmi <weight> <kg|lbs> <height> <cm|in>\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bmi --tui   # Run in Text-based UI mode\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bmi --web   # Run in Web mode\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bmi --version  # Show version\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\nExamples:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bmi 70 kg 170 cm\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bmi 154 lbs 67 in\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bmi --tui\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  bmi --web\n")
		flag.PrintDefaults()
	}

	// Flags
	showVersion := flag.Bool("version", false, "Display the version")
	useTUI := flag.Bool("tui", false, "Run in TUI mode")
	useWeb := flag.Bool("web", false, "Run in Web mode")
	flag.Parse()

	// Show version and exit if --version is provided
	if *showVersion {
		log.Printf("Version requested. BMI Calculator version: %s", version)
		log.WithField("version", version).Info("Version requested")
		fmt.Println("BMI Calculator version:", version)
		os.Exit(0)
	}

	// Start Web Server if --web is provided
	if *useWeb {
		log.Info("Starting web server on http://localhost:8080")
		web.StartWeb()
		return
	}

	// Run TUI if --tui is provided
	if *useTUI {
		log.Info("Starting TUI mode")
		err := ui.StartTUI()
		if err != nil {
			log.WithError(err).Fatal("Error running TUI")
		}
		return
	}

	// CLI mode expects 4 arguments: <weight> <kg|lbs> <height> <cm|in>
	if len(os.Args) != 5 {
		log.Error("Invalid number of arguments for CLI mode")
		fmt.Fprintln(os.Stderr, "Error: Invalid number of arguments.")
		flag.Usage()
		os.Exit(1)
	}

	log.Info("Starting CLI mode")
	weightStr, weightUnit, heightStr, heightUnit := os.Args[1], os.Args[2], os.Args[3], os.Args[4]

	// Parse inputs
	weight, height, err := input.ParseInput(weightStr, heightStr, weightUnit, heightUnit)
	if err != nil {
		log.WithError(err).Fatal("Input parsing failed")
	}

	// Calculate BMI
	bmiValue := bmi.CalculateBMI(weight, height)
	result := bmi.InterpretBMI(bmiValue)

	// Output results
	log.Info("BMI calculated successfully")
	fmt.Printf("Your BMI: %.2f\nCategory: %s\n", bmiValue, result)
}
