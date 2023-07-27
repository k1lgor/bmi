/*
Copyright © 2023 Plamen Ivanov <paco.iwanow@gmail.com>
*/
package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// apiCmd represents the api endpoint command
var (
	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "API endpoint",
		Long:  "API endpoint",
		Run: func(cmd *cobra.Command, args []string) {
			// Create a router
			r := mux.NewRouter()

			// Add an HTML
			r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `
					<html>
					<head>
						<title>BMI Calculator</title>
					</head>
					<style>
						body {
							font-family: sans-serif;
							margin: 20px auto;
							padding: 0;
						}

						form {
							width: 50%;
							margin: 20px auto;
						}

						input {
							width: 50%;
							border: 1px solid #ccc;
							padding: 10px;
							margin: 20px auto;
							display: block;
						}

						button {
							background-color: #000;
							color: #fff;
							border: none;
							padding: 10px;
							margin: 10px 0;
							cursor: pointer;
							display: block;
						}

					</style>
					<body>
						<form action="/bmi" method="post">
							<input type="text" name="height" placeholder="Height (in cm)">
							<input type="text" name="weight" placeholder="Weight (in kg)">
							<input type="submit" value="Calculate BMI">
						</form>
						<div class="result"></div>
						<div class="status"></div>
					</body>
					</html>
				`)
			})

			// Add an endpoint for BMI
			r.HandleFunc("/bmi", func(w http.ResponseWriter, r *http.Request) {
				// Get the height and weight from the request
				height, err := strconv.ParseFloat(r.FormValue("height"), 64)
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintf(w, "Invalid height")
					return
				}

				weight, err := strconv.ParseFloat(r.FormValue("weight"), 64)
				if err != nil {
					w.WriteHeader(400)
					fmt.Fprintf(w, "Invalid weight")
					return
				}

				bmi := weight / (height / 100) / (height / 100)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(200)
				fmt.Fprintf(w, `{"bmi": %.2f}`, bmi)
			})
			// Start the server
			http.ListenAndServe(":8080", r)
		},
	}
)

func init() {
	rootCmd.AddCommand(apiCmd)
}
