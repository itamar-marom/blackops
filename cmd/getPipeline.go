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

// getPipelineCmd represents the getPipeline command
var getPipelineCmd = &cobra.Command{
	Use:   "pipeline <name>",
	Short: "Get a pipeline",
	Long:  `Get a pipline definition from a specific application and a provider`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getPipeline called")
	},
}

func init() {
	getPipelineCmd.PersistentFlags().StringP("application", "a", "", "application to be part of")
	getPipelineCmd.PersistentFlags().StringP("provider", "p", "", "provider to create from")

	getPipelineCmd.MarkPersistentFlagRequired("application")
	getPipelineCmd.MarkPersistentFlagRequired("provider")

	getCmd.AddCommand(getPipelineCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getPipelineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getPipelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
