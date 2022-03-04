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

// repositoryCmd represents the repository command
var repositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Creates a repository definition",
	Long:  `Creates a repository definition in cache -> ~/.blackops`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("repository called")
		fmt.Println("Here are the arguments of card command : " + strings.Join(args, ","))
	},
}

func init() {
	getCmd.AddCommand(repositoryCmd)
	createCmd.AddCommand(repositoryCmd)
	deleteCmd.AddCommand(repositoryCmd)

	repositoryCmd.PersistentFlags().StringP("url", "u", "", "URL of the repository")
	repositoryCmd.PersistentFlags().StringP("token", "t", "", "Token for the repository")

	repositoryCmd.MarkPersistentFlagRequired("url")
	repositoryCmd.MarkPersistentFlagRequired("token")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
