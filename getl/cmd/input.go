// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"strings"

	"github.com/ibbd-dev/getl/common"
	"github.com/ibbd-dev/getl/input"
	"github.com/spf13/cobra"
)

type InputParams struct {
	Type     string
	Filename string
}

var inputParams InputParams
var supportInputTypes = strings.Join(input.SupportTypes, ",")

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "输入数据",
	Long: `支持从各种格式的数据输入，暂时支持的格式：
- csv

计划支持的格式：

- text
- mysql
`,
	Run: func(cmd *cobra.Command, args []string) {
		var is_in bool
		for _, v := range input.SupportTypes {
			if v == inputParams.Type {
				is_in = true
			}
		}
		if !is_in {
			panic("非法的输入类型")
		}
	},
}

func init() {
	rootCmd.AddCommand(inputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	inputCmd.PersistentFlags().StringVar(&inputParams.Filename, "filename", "", "文件名")
	inputCmd.PersistentFlags().StringVar(&inputParams.Type, "type", common.TypeCSV, "输入类型，支持的类型："+supportInputTypes)
}
