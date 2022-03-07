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

// createApplicationCmd represents the createApplication command
var createApplicationCmd = &cobra.Command{
	Use:   "application <name>",
	Short: "Create an application",
	Long:  `Create an application definition in a given repository`,
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		repositoryName, _ := cmd.Flags().GetString("repository")
		propertiesFilePath, _ := cmd.Flags().GetString("properties")
		repository, _, err := models.GetRepository(repositoryName)
		utils.CheckError(err)
		models.CreateApplication(appName, repository, propertiesFilePath)
	},
}

func init() {
	createApplicationCmd.PersistentFlags().StringP("repository", "r", "", "repository for application ecosystem")
	createApplicationCmd.PersistentFlags().StringP("properties", "p", "", "properties file for this application")

	createApplicationCmd.MarkPersistentFlagRequired("repository")
	createApplicationCmd.MarkPersistentFlagRequired("properties")

	createCmd.AddCommand(createApplicationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createApplicationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createApplicationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
