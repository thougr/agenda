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

// MeetingQueryCmd represents the MeetingQuery command
var MeetingQueryCmd = &cobra.Command{
	Use:   "queryM -s [StartTime] -e [EndTime]",
	Short: "To query all the meeting have attended during [StartTime] and [EndTime]",
	Long: `You can query all the meeting have attended during [StartTime] and [EndTime]`,
	Run: func(cmd *cobra.Command, args []string) {
		debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
		if entity.StartAgenda() == false {
			debugLog.Println("Fail, please log in")
			fmt.Println("Fail, please log in")
		}
		arg_s, _ := cmd.Flags().GetString("StartTime")
		arg_e, _ := cmd.Flags().GetString("EndTime")

		mm := entity.MeetingQuery(entity.CurrentUser.Name, arg_s, arg_e)
		if len(mm)!= 0 {
			debugLog.Println("Query meeting successfully")
			fmt.Println("Query meeting successfully\n")
			fmt.Println("Sponsor Title StartDate EndDate Participators")
			for i, m := range mm {
				fmt.Printf("%d. %s\n", i+1, m)
			}
			
		} else {
			debugLog.Println("Fail to query meeting")
			fmt.Println("Fail to query meeting")
		}
		entity.QuitAgenda()
	},
}

func init() {
	RootCmd.AddCommand(MeetingQueryCmd)
	MeetingQueryCmd.Flags().StringP("StartTime", "s", "", "meeting's startTime")
	MeetingQueryCmd.Flags().StringP("EndTime", "e", "", "meeting's endTime")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MeetingQueryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MeetingQueryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
