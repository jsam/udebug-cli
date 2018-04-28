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
	"github.com/spf13/cobra"
	"github.com/jsam/udebug-cli/utils"
	"fmt"
	"os"
	"text/tabwriter"
	"github.com/jsam/udebug-cli/client/output"
)

func getOutput(inputID string) string {
	client := utils.NewAPIClient()

	params := output.NewGetOutputAPIOutputRetrieveJSONParams()
	params.InputID = inputID

	resp, err := client.Client.Output.GetOutputAPIOutputRetrieveJSON(params, client.Auth)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(resp.Payload) != 1 {
		panic("Please report this issue at https://github.com/jsam/udebug-cli")
	}

	return resp.Payload[0]
}

// outputCmd represents the output command
var outputCmd = &cobra.Command{
	Use:   "output",
	Short: "Fetches the output data for your program",
	Long: `You must specify  input identifier as positional arguments. For example:

	udebug-cli fetch output 809839`,
	Run: func(cmd *cobra.Command, args []string) {
		inputID := args[0]
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', tabwriter.TabIndent)

		outputs := getOutput(inputID)
		line := fmt.Sprintf("%s", outputs)
		fmt.Fprintln(w, line)
		fmt.Println(len(outputs))
		w.Flush()
	},
}

func init() {
	fetchCmd.AddCommand(outputCmd)
}
