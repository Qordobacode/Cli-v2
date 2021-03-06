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

package segment

import (
	"fmt"
	"github.com/qordobacode/cli-v2/pkg/general/log"
	"github.com/qordobacode/cli-v2/pkg/types"
	"github.com/spf13/cobra"
)

// updateValueCmd represents the updateValue command
var (
	updateKeyVersion string
	updateKeyKey     string
	updateKeyValue   string
	updateKeyRef     string
)

// NewUpdateSegmentCommand function add `update-value` command
func NewUpdateSegmentCommand() *cobra.Command {
	updateValueCmd := &cobra.Command{
		Annotations: map[string]string{"group": "segment"},
		Use:         "update-value",
		Short:       "Update value by key",
		Example:     `qor update-value file_name.doc --version v1 --key "/go_nav_menu" --value "text text text" --ref "Main nav text"`,
		PreRunE:     preValidateUpdateKeyParameters,
		Run:         updateValue,
	}

	updateValueCmd.Flags().StringVarP(&updateKeyVersion, "version", "v", "", "file version")
	updateValueCmd.Flags().StringVarP(&updateKeyKey, "key", "k", "", "key to add")
	updateValueCmd.Flags().StringVar(&updateKeyValue, "value", "", "value to add")
	updateValueCmd.Flags().StringVarP(&updateKeyRef, "ref", "r", "", "")
	return updateValueCmd
}

func preValidateUpdateKeyParameters(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("filename is mandatory")
	}
	if updateKeyKey == "" {
		return fmt.Errorf("flag 'key' is mandatory")
	}
	if updateKeyValue == "" {
		return fmt.Errorf("flag 'value' is mandatory")
	}
	startLocalServices(cmd, args)
	return nil
}

func updateValue(cmd *cobra.Command, args []string) {
	if appConfig == nil {
		log.Errorf("error occurred on configuration load: ")
		return
	}
	keyAddRequest := &types.KeyAddRequest{
		Key:       updateKeyKey,
		Source:    updateKeyValue,
		Reference: updateKeyRef,
	}
	segmentService.UpdateKey(args[0], updateKeyVersion, keyAddRequest)
}
