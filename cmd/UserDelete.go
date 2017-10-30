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

// UserDeleteCmd represents the UserDelete command
var UserDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete your account in Agenda",
	Long: `you can delete your account in the database of Agenda:

attention:After deleting,you will need to register a new User to login Agenda.`,
Run: func(cmd *cobra.Command, args []string) {
	debugLog := log.New(logFile,"[Result]", log.Ldate|log.Ltime|log.Lshortfile)
	if entity.StartAgenda() == false {
		debugLog.Println("Fail, please log in")
		fmt.Println("Fail, please log in")
	}

	if entity.DeleteUser(entity.CurrentUser.Name, entity.CurrentUser.Password) {
		debugLog.Println("Delete this account successfully")
		fmt.Println("Delete this account successfully")
	} else {
		debugLog.Println("Fail to delete this account")
		fmt.Println("Fail to delete this account")
	}
	entity.QuitAgenda()
},
}

func init() {
	RootCmd.AddCommand(UserDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UserDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UserDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
