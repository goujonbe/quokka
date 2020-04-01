/*
Copyright © 2020 Benoît Goujon

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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get config parameters",
	Example: "quokka get password-length",
	Run: func(cmd *cobra.Command, args []string) {
		configParameterWanted := args[0]
		switch configParameterWanted {
		case "password-length":
			length := viper.GetInt("password.length")
			if length != 0 {
				fmt.Println("password length: ", length)
			} else {
				fmt.Println("Password length not found. Either config file does not exist or password length is not set.")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
