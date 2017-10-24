// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// UserQueryCmd represents the UserQuery command
var UserQueryCmd = &cobra.Command{
	Use:   "queryU",
	Short: "To query all the users' names",
	Long: `You can query all the users's names who have registed.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UserQuery called")
	},
}

func init() {
	RootCmd.AddCommand(UserQueryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UserQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UserQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
