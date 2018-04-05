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

	//"github.com/ibbd-dev/getl/io"
	"github.com/spf13/cobra"
)

// similarCmd represents the similar command
var similarCmd = &cobra.Command{
	Use:   "similar",
	Short: "计算文本相似性",
	Long: `对于输入的文本计算相似性
`,
	Example: ``,
	Run: func(cmd *cobra.Command, args []string) {
		/*
			data, err := params.ReadAll()
			if err != nil {
				panic(err)
			}

			params.Write(data)
		*/
		fmt.Println("similar called")
	},
}

func init() {
	rootCmd.AddCommand(similarCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// similarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// similarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
