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
	"github.com/spf13/cobra"
)

// getRepositoryCmd represents the getRepository command
var getRepositoryCmd = &cobra.Command{
	Use:   "repository <name>",
	Short: "Get a repository definition",
	Long:  `Get a repository definition in cache -> ~/.blackops`,
	Run: func(cmd *cobra.Command, args []string) {
		repository, err := models.GetRepository(args[0])
		if err != nil {
			println(err.Error())
		} else {
			println("Name: " + repository.Name)
			println("URL: " + repository.URL)
		}
	},
}

func init() {
	getCmd.AddCommand(getRepositoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getRepositoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getRepositoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
