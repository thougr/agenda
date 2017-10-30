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
	"agenda/entity"
	"fmt"
	"log"
	"github.com/spf13/cobra"
)

// MeetingsClearCmd represents the MeetingsClear command
var MeetingsClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "clear all the meeting created by the current user",
	Long: `you can clear all the meeting you have created`,
	Run: func(cmd *cobra.Command, args []string) {
		debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.StartAgenda() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in")
		}

		if entity.DeleteAllMeetings(entity.CurrentUser.Name) {
			debugLog.Println("Clear meeting successfully")
			fmt.Println("Clear meeting successfully")
			
		} else {
			debugLog.Println("Fail to clear meeting")
			fmt.Println("Fail to clear meeting")
		}
		entity.QuitAgenda()
	},
}

func init() {
	RootCmd.AddCommand(MeetingsClearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MeetingsClearCmd.PersistentFlags().String("foo", "", "A help for foo")
MeetingsClearCmd.Flags().StringP("username", "u", "", "new user's username")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MeetingsClearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
