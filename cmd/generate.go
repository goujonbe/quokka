/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random password",
	Long: `
Generate password according to configuration
parameters in .quokka config file or default parameters
if config file does not exist.`,
	Example: `
quokka generate password
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println("Please clarify what you want to generate. Possible args: password")
			return nil
		}

		if len(args) != 1 {
			fmt.Errorf("You can only generate one thing at the same time.")
		}

		thingToGenerate := args[0]
		switch thingToGenerate {
		case "password":
			res, err := password.Generate(12, 5, 5, false, false)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Random password: ", res)
		default:
			return fmt.Errorf("%s generation not implemented yet. Feel free to open an issue asking for this feature.", thingToGenerate)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
