/*
Copyright © 2020 Benoit Goujon

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

	password "github.com/goujonbe/quokka/password"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random password",
	Long: `
Generate password according to configuration
parameters in .quokka config file or default parameters
if config file does not exist.`,
	Example:   "quokka generate password",
	ValidArgs: []string{"password"},
	Args:      cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		thingToGenerate := args[0]
		switch thingToGenerate {
		case "password":
			var length int
			if viper.IsSet("password.length") && !cmd.Flags().Changed("length") {
				fmt.Println("Using password length in config file.")
				length = viper.GetInt("password.length")
			} else if cmd.Flags().Changed("length") || !viper.IsSet("password.length") {
				// get flag value or default one
				lengthInFlag, err := cmd.Flags().GetInt("length")
				length = lengthInFlag
				if err != nil {
					return fmt.Errorf("Could not parse password length: %v", err)
				}
			}

			res, err := password.Generate(length, true, true, true, true)
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
	generateCmd.Flags().IntP("length", "l", 12, "Password length")
}
