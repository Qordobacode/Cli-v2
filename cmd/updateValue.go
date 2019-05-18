// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
)

// updateValueCmd represents the updateValue command
var (
	updateValueCmd = &cobra.Command{
		Use:   "updateValue",
		Short: "Update value by key",
		Run:   updateValue,
	}
	updateKeyVersion string
	updateKeyKey     string
	updateKeyValue   string
	updateKeyRef     string
)

func init() {
	rootCmd.AddCommand(updateValueCmd)
	updateValueCmd.Flags().StringVarP(&updateKeyVersion, "version", "v", "", "file version")
	updateValueCmd.Flags().StringVarP(&updateKeyKey, "key", "k", "", "key to add")
	updateValueCmd.Flags().StringVar(&updateKeyValue, "value", "", "value to add")
	updateValueCmd.Flags().StringVarP(&updateKeyRef, "ref", "r", "", "")
}

func updateValue(cmd *cobra.Command, args []string) {

}