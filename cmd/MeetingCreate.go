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

// MeetingCreateCmd represents the MeetingCreate command
var MeetingCreateCmd = &cobra.Command{
	Use:   "create -t [Title] -p [Participator] -s [StartTime] -e [EndTime]",
	Short: "To create a new meeting",
	Long: `To create a new meeting with:

[Title] the Title of the meeting
[Participator] the Participator of the meeting,the Participator can only attend one meeting during one meeting time
[StartTime] the StartTime of the meeting
[EndTime] the EndTime of the meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.StartAgenda() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in")
		}
		arg_t, _ := cmd.Flags().GetString("Title")
		arg_p, _ := cmd.Flags().GetStringSlice("Participator")
		arg_s, _ := cmd.Flags().GetString("StartTime")
		arg_e, _ := cmd.Flags().GetString("EndTime")
		if entity.CreateMeeting(entity.CurrentUser.Name, arg_t, arg_s, arg_e, arg_p) {
			debugLog.Println("Create meeting successfully")
			fmt.Println("Create meeting successfully")
		} else {
			debugLog.Println("Fail to create meeting")
			fmt.Println("Fail to create meeting")
		}
		entity.QuitAgenda()
	},
}

func init() {
	RootCmd.AddCommand(MeetingCreateCmd)
	MeetingCreateCmd.Flags().StringP("Title", "t", "", "meeting title")
	MeetingCreateCmd.Flags().StringSliceP("Participator", "p", []string{}, "meeting's participator")
	MeetingCreateCmd.Flags().StringP("StartTime", "s", "", "meeting's startTime")
	MeetingCreateCmd.Flags().StringP("EndTime", "e", "", "meeting's endTime")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MeetingCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MeetingCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
