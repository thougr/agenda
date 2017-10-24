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

// RemoveParticipatorCmd represents the RemoveParticipator command
var RemoveParticipatorCmd = &cobra.Command{
	Use:   "remove -p [Participator] -t [Title]",
	Short: "To remove Participator from the meeting",
	Long: `remove [Participator] from the meeting with the title of [Title]:

attention:If there is no Participators in the meeting,the meeting will be deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RemoveParticipator called")
	},
}

func init() {
	RootCmd.AddCommand(RemoveParticipatorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RemoveParticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RemoveParticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
