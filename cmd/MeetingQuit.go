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

// MeetingQuitCmd represents the MeetingQuit command
var MeetingQuitCmd = &cobra.Command{
	Use:   "quit -t [title]",
	Short: "quit the meeting with the title [title]",
	Long: `you can quit the meeting with the title of [title]:

attention:if there is no participators in this meeting,the meeting will be deleted`,
	Run: func(cmd *cobra.Command, args []string) {
		debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.StartAgenda() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in")
		}
		arg_t, _ := cmd.Flags().GetString("Title")

		if entity.DeleteMeeting(entity.CurrentUser.Name, arg_t) {
			debugLog.Println("Quit meeting successfully")
			fmt.Println("Quit meeting successfully")
		} else {
			debugLog.Println("Fail to quit meeting")
			fmt.Println("不存在该会议或者该会议不是本用户创建")
			fmt.Println("Fail to quit meeting")
		}
		entity.QuitAgenda()
	},
}

func init() {
	RootCmd.AddCommand(MeetingQuitCmd)
	MeetingQuitCmd.Flags().StringP("Title", "t", "", "meeting title")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MeetingQuitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MeetingQuitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
