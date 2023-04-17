/*
Copyright © 2023 Plamen Ivanov <paco.iwanow@gmail.com>
*/
package cmd

import (
	"fmt"
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
		if len(args) != 2 {
			fmt.Printf("2 paramaters are required, e.g. bmi 60 170\n")
			os.Exit(1)
		} else {
			fmt.Println("Your BMI is being calculated 🖩")
		}

		weight, err := strconv.Atoi(os.Args[1])
		check(err)

		height, err := strconv.Atoi(os.Args[2])
		check(err)
		total := BMI(weight, height)

		if total < 18.5 {
			fmt.Printf("Your BMI is: %.2f\nYou are underweight.", total)
		} else if 18.5 <= total && total < 24.9 {
			fmt.Printf("Your BMI is: %.2f\nYou are in normal weight.", total)
		} else if 25 <= total && total < 29.9 {
			fmt.Printf("Your BMI is: %.2f\nYou are overweight.", total)
		} else {
			fmt.Printf("Your BMI is: %.2f\nYou are obese.", total)
		}
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
		fmt.Println(err)
		os.Exit(1)
	}
}

func BMI(w, h int) float64 {
	return float64(w) / math.Pow(float64(h)/100, 2)
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
