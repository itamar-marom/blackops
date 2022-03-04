/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// createRepositoryCmd represents the createRepository command
var createRepositoryCmd = &cobra.Command{
	Use:   "repository <name>",
	Short: "Create a repository definition",
	Long:  `Create a repository definition in cache -> ~/.blackops`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createRepository called")
		fmt.Println("Here are the arguments of card command: " + strings.Join(args, ","))
	},
}

func init() {
	createRepositoryCmd.PersistentFlags().StringP("url", "u", "", "URL of the repository")
	createRepositoryCmd.PersistentFlags().StringP("token", "t", "", "Token for the repository")

	createRepositoryCmd.MarkPersistentFlagRequired("url")
	createRepositoryCmd.MarkPersistentFlagRequired("token")

	createCmd.AddCommand(createRepositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createRepositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createRepositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
