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

	"microservice-agenda/cli/entity"
	"microservice-agenda/cli/errors"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "for guest to login",
	Long:  `for guests to enter correct username and password to login `,
	// Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		if username == "" {
			errors.ErrorMsg("login error", "username required!")
		}
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			errors.ErrorMsg("login error", "password required!")
		}

		if entity.Login(username, password) {
			fmt.Println("user:" + username + " login successfully")
		} else {
			fmt.Println("failed to log in!")
		}
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout",
	Long:  `log out as a guest`,
	// Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cuName, got := entity.GetCurrentUser()
		if !got {
			return
		}
		entity.Logout(cuName)
		fmt.Println("logout successfully")
	},
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "to register a new user",
	Long: `register a new user with this command ,followed by a unique username and
	password. If the user by the given username is already existed, the register
	action will fail.`,
	// Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		if username == "" {
			errors.ErrorMsg("register error", "username required.")
		}
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			errors.ErrorMsg("register error", "password required!")
		}

		if entity.Register(username, password) {
			fmt.Println("user:" + username + " register successfully!")
		} else {
			fmt.Println("fail to register")
		}
	},
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "list all users",
	Long:  `list all users' information`,
	// Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cuName, got := entity.GetCurrentUser()
		if !got {
			return
		}
		entity.LookupAllUser(cuName)
	},
}

var cancelUserCmd = &cobra.Command{
	Use:   "cancelUser",
	Short: "remove an account from users",
	Long:  `remove an account from users`,
	// Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cuName, got := entity.GetCurrentUser()
		if !got {
			return
		}
		if entity.CancelAccount(cuName) {
			fmt.Println("cancel " + cuName + " account successfully")
		} else {
			fmt.Println("fail to cancel " + cuName + " account")
		}
	},
}

var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting",
	Short: "quit from all meetings with you as participator",
	Long:  `quit from all meetings with you as participator`,
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cuName, got := entity.GetCurrentUser()
		if !got {
			return
		}
		title, _ := cmd.Flags().GetString("title")
		if entity.QuitMeeting(title, cuName) {
			fmt.Println("quit meeting " + title + " successfully")
		} else {
			fmt.Println("fail to quit meeting " + title)
		}
	},
}

var clearMeetingsCmd = &cobra.Command{
	Use:   "clearMeetings",
	Short: "clear all meetings with you as sponsor",
	Long:  `clear all meetings with you as sponsor`,
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cuName, got := entity.GetCurrentUser()
		if !got {
			return
		}
		entity.ClearAllMeetings(cuName)
	},
}

var cancelMeetingsCmd = &cobra.Command{
	Use:   "cancelMeeting",
	Short: "cancel meetings you sponsored with specified title",
	Long:  `cancel meetings you sponsored with specified title`,
	Run: func(cmd *cobra.Command, args []string) {
		cuName, got := entity.GetCurrentUser()
		if !got {
			return
		}
		title, _ := cmd.Flags().GetString("title")
		entity.CancelMeeting(title, cuName)
	},
}

var setEmailCmd = &cobra.Command{
	Use:   "setEmail",
	Short: "set registered user's email",
	Long:  `set registered user's email`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cuName, got := entity.GetCurrentUser()
		if !got {
			return
		}
		email, _ := cmd.Flags().GetString("email")
		if email == "" {
			errors.ErrorMsg(cuName, "valid email required.")
			return
		}
		entity.SetEmail(email, cuName)
		fmt.Println("set " + cuName + "'s email to be " + email)
	},
}

var setTelCmd = &cobra.Command{
	Use:   "setTel",
	Short: "set registered user's telephone number",
	Long:  `set registered user's telephone number`,
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cuName, got := entity.GetCurrentUser()
		if !got {
			return
		}
		tel, _ := cmd.Flags().GetString("telephone")
		if tel == "" {
			errors.ErrorMsg(cuName, "valid telephone number required.")
			return
		}
		entity.SetTelephone(tel, cuName)
		fmt.Println("set " + cuName + "'s telephone to be " + tel)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "the name of user to log in")
	loginCmd.Flags().StringP("password", "p", "", "the password of user to log in")

	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "the name of new user to be created")
	registerCmd.Flags().StringP("password", "p", "", "the password of user to be created")

	quitMeetingCmd.Flags().StringP("title", "t", "", "the title of the meeting to quit")
	setEmailCmd.Flags().StringP("email", "e", "", "set current user's email")
	setTelCmd.Flags().StringP("telephone", "t", "", "set current user's telephone number")

	cancelMeetingsCmd.Flags().StringP("title", "t", "", "title of the meeting to be canceled")

	RootCmd.AddCommand(setTelCmd)
	RootCmd.AddCommand(setEmailCmd)
	RootCmd.AddCommand(logoutCmd)
	RootCmd.AddCommand(usersCmd)
	RootCmd.AddCommand(cancelUserCmd)
	RootCmd.AddCommand(quitMeetingCmd)
	RootCmd.AddCommand(clearMeetingsCmd)
	RootCmd.AddCommand(cancelMeetingsCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCommandsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCommandsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
