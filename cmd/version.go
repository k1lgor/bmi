/*
Copyright © 2023 Plamen Ivanov <paco.iwanow@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var (
	version    = "0.2.0"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version number of bmi",
		Long:  "Prints the version number of bmi",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("bmi current version: %s\n", version)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
