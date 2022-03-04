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

	"github.com/spf13/cobra"
)

// pipelineCmd represents the pipeline command
var pipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Creates a pipeline",
	Long:  `Creates a pipline definition for a specific application in a given repository`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pipeline called")
	},
}

func init() {
	getCmd.AddCommand(pipelineCmd)
	createCmd.AddCommand(pipelineCmd)
	deleteCmd.AddCommand(pipelineCmd)

	pipelineCmd.PersistentFlags().StringP("application", "a", "", "application to be part of")
	pipelineCmd.PersistentFlags().StringP("api-resource", "r", "", "api-resource to create from")

	pipelineCmd.MarkPersistentFlagRequired("application")
	pipelineCmd.MarkPersistentFlagRequired("api-resource")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pipelineCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pipelineCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
