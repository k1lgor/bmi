/*
Copyright © 2023 Plamen Ivanov <paco.iwanow@gmail.com>
*/
package cmd

import (
	"github.com/apsdehal/go-logger"
	"github.com/spf13/cobra"
	"math"
	"os"
	"strconv"
)

var rootCmd = &cobra.Command{
	Use:   "bmi <weight> <height>",
	Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
	Short: "BMI (Body Mass Index) is a measure of body fat based on a person's weight and height.",
	Long: `BMI (Body Mass Index) is a measure of body fat based on a person's weight and height.

The CLI takes 2 integer arguments.
First one for your weight in kg
Second one for height in cm`,

	Run: func(cmd *cobra.Command, args []string) {
		weight, err := strconv.Atoi(os.Args[1])
		check(err)

		height, err := strconv.Atoi(os.Args[2])
		check(err)

		total, weightCategory := BMI(weight, height)
		loggerOutput(total, weightCategory)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func loggerOutput(total float64, weightCategory string) {
	log, err := logger.New("bmi", 1, os.Stdout)
	check(err)

	log.SetFormat("%{message}")
	log.Info("Your BMI is being calculated 🖩")
	switch weightCategory {
	case "underweight":
		log.WarningF("Your BMI is: %.2f\nYour weight is: %s", total, weightCategory)
	case "normal":
		log.NoticeF("Your BMI is: %.2f\nYour weight is: %s", total, weightCategory)
	case "overweight":
		log.WarningF("Your BMI is: %.2f\nYour weight is: %s", total, weightCategory)
	case "obese":
		log.ErrorF("Your BMI is: %.2f\nYour weight is: %s", total, weightCategory)
	}
	return
}

func BMI(w, h int) (total float64, weightCategory string) {
	total = float64(w) / math.Pow(float64(h)/100, 2)

	if total < 18.5 {
		weightCategory = "underweight"
	} else if 18.5 <= total && total < 24.9 {
		weightCategory = "normal"
	} else if 25 <= total && total < 29.9 {
		weightCategory = "overweight"
	} else {
		weightCategory = "obese"
	}
	return total, weightCategory
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bmi.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
