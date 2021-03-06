/*
Copyright © 2022 Itamar marom <marom.itamar@gmail.com>

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
	"github.com/itamar-marom/blackops/models"
	"github.com/itamar-marom/blackops/utils"
	"github.com/spf13/cobra"
)

// getApplicationCmd represents the getApplication command
var getApplicationCmd = &cobra.Command{
	Use:   "application <name>",
	Short: "Get an application",
	Long:  `Get an application definition from a specific repository`,
	Run: func(cmd *cobra.Command, args []string) {

		repoName, _ := cmd.Flags().GetString("repository")
		r, _, err := models.GetRepository(repoName)
		utils.CheckError(err)

		if len(args) == 0 {
			models.PrintApplications(r)
		} else {
			models.PrintApplicationByName(args[0], r)
		}
	},
}

func init() {
	getApplicationCmd.PersistentFlags().StringP("repository", "r", "", "repository for application ecosystem")

	getApplicationCmd.MarkPersistentFlagRequired("repository")

	getCmd.AddCommand(getApplicationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getApplicationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getApplicationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
