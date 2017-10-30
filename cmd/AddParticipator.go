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
	"log"
    "agenda/entity"
	"github.com/spf13/cobra"
)

// AddParticipatorCmd represents the AddParticipator command
var AddParticipatorCmd = &cobra.Command{
	Use:   "add -p [Participator] -t [Title]",
	Short: "To add Participator of the meeting",
	Long: `Add [Participator] to the meeting with the title of [Title]:

attention:If the Participator cannot attend during the time, add fail.`,
Run: func(cmd *cobra.Command, args []string) {
	debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
	if entity.StartAgenda() == false {
		debugLog.Println("Fail,please log in")
		fmt.Println("Fail,please log in")
	}
	
	arg_p, _ := cmd.Flags().GetStringSlice("Participator")
	arg_t, _ := cmd.Flags().GetString("Title")
	if entity.Addparticipator(arg_t, arg_p) {
		debugLog.Println("Add participators successfully")
		fmt.Println("Add participators successfully")
	} else {
		debugLog.Println("Fail to add participators")
		fmt.Println("Fail to add participators")
	}
	entity.QuitAgenda()
  },
}

func init() {
	RootCmd.AddCommand(AddParticipatorCmd)
	AddParticipatorCmd.Flags().StringSliceP("Participator", "p", []string{}, "meeting's participator")
	AddParticipatorCmd.Flags().StringP("Title", "t", "", "meeting title")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// AddParticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// AddParticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
