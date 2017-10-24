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

/*注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；
成功注册后，亦应反馈一个成功注册的信息。*/
// UserRegisterCmd represents the UserRegister command
var UserRegisterCmd = &cobra.Command{
	Use:   "register -u [UserName] -p [Pass] -e [Email] -t [Phone]",
	Short: "Register a new user",
	Long: `Input command register -u UserName -p PassWord -e Email -t Phone:

[Username] is the name of the new register
[PassWord] is the password to login
[Email]is the email address of the register
[Phone] is the phone of the register`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UserRegister called")
	},
}

func init() {
	RootCmd.AddCommand(UserRegisterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UserRegisterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UserRegisterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
