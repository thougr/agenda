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

// MeetingQueryCmd represents the MeetingQuery command
var MeetingQueryCmd = &cobra.Command{
	Use:   "queryM -t [StartTime] -e [EndTime]",
	Short: "To query all the meeting have attended during [StartTime] and [EndTime]",
	Long: `You can query all the meeting have attended during [StartTime] and [EndTime]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MeetingQuery called")
	},
}

func init() {
	RootCmd.AddCommand(MeetingQueryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MeetingQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MeetingQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
