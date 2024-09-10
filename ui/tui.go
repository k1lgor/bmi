package ui

import (
	"fmt"

	"bmi/pkg/bmi"
	"bmi/pkg/input"

	"github.com/rivo/tview"
)

func StartTUI() error {
	app := tview.NewApplication()
	form := tview.NewForm()

	var weightStr, heightStr, weightUnit, heightUnit string

	form.AddInputField("Weight", "", 20, nil, func(text string) {
		weightStr = text
	}).
		AddDropDown("Weight Unit", []string{"kg", "lbs"}, 0, func(option string, index int) {
			weightUnit = option
		}).
		AddInputField("Height", "", 20, nil, func(text string) {
			heightStr = text
		}).
		AddDropDown("Height Unit", []string{"cm", "in"}, 0, func(option string, index int) {
			heightUnit = option
		}).
		AddButton("Calculate", func() {
			weight, height, err := input.ParseInput(weightStr, heightStr, weightUnit, heightUnit)
			if err != nil {
				modal := tview.NewModal().SetText(fmt.Sprintf("Error: %v", err)).AddButtons([]string{"OK"})
				app.SetRoot(modal, true)
				return
			}

			// Calculate BMI
			bmiValue := bmi.CalculateBMI(weight, height)
			result := bmi.InterpretBMI(bmiValue)
			modal := tview.NewModal().SetText(fmt.Sprintf("Your BMI: %.2f\n%s", bmiValue, result)).AddButtons([]string{"OK"})
			app.SetRoot(modal, true)
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("BMI Calculator")
	return app.SetRoot(form, true).Run()
}
