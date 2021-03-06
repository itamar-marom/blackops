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
	"fmt"

	"github.com/spf13/cobra"
)

// createPipelineCmd represents the createPipeline command
var createPipelineCmd = &cobra.Command{
	Use:   "pipeline <name>",
	Short: "Create a pipeline",
	Long:  `Create a pipline definition for a specific application and a provider`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createPipeline called")
	},
}

func init() {
	createPipelineCmd.PersistentFlags().StringP("application", "a", "", "application to be part of")
	createPipelineCmd.PersistentFlags().StringP("provider", "p", "", "provider to create from")

	createPipelineCmd.MarkPersistentFlagRequired("application")
	createPipelineCmd.MarkPersistentFlagRequired("provider")

	createCmd.AddCommand(createPipelineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createPipelineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createPipelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
