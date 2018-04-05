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
	"io"
	"os"
	"strings"

	"github.com/ibbd-dev/getl/common"
	"github.com/ibbd-dev/getl/output"
	"github.com/ibbd-dev/getl/tools/csv"
	"github.com/spf13/cobra"
)

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "数据输出",
	Long: `将数据输出到相应媒介, 支持的格式：
- csv

后续计划支持的：

- stdout
- mysql
`,
	Run: func(cmd *cobra.Command, args []string) {
		var is_in bool
		for _, v := range output.SupportTypes {
			if v == inputParams.Type {
				is_in = true
			}
		}
		if !is_in {
			panic("非法的输入类型")
		}

		var err error
		if inputParams.Type == common.TypeCSV {
			if inputParams.Filename == "" {
				panic("文件名不能为空")
			}

			err = outputCSV()
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(outputCmd)

	inputCmd.PersistentFlags().StringVar(&inputParams.Filename, "filename", "", "文件名")
	inputCmd.PersistentFlags().StringVar(&inputParams.Type, "type", common.TypeCSV, "输出类型，支持的类型："+supportInputTypes)
}

func outputCSV() (err error) {
	var fieldnames []string
	reader := common.NewPipelineIO()
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Read Error: ", err, "   Filename: ", Filename)
			return
		}

		if len(fieldnames) == 0 {
			for field, _ := range record {
				fieldnames = append(fieldnames, field)
			}
		}

	}

	return err
}
