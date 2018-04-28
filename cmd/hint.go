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
	"github.com/jsam/udebug-cli/client/hint"
)

// hintCmd represents the hint command
var hintCmd = &cobra.Command{
	Use:   "hint",
	Short: "Fetches the specified hint",
	Long: `You must specify hint identifier as positional arguments. For example:

	udebug-cli fetch hint 809839`,
	Run: func(cmd *cobra.Command, args []string) {
		hintID := args[0]
		client := utils.NewAPIClient()

		params := hint.NewGetHintAPIHintRetrieveJSONParams()
		params.HintID = hintID

		resp, err := client.Client.Hint.GetHintAPIHintRetrieveJSON(params, client.Auth)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', tabwriter.TabIndent)

		for _, input := range resp.Payload {
			line := fmt.Sprintf("%s", input)
			fmt.Fprintln(w, line)
		}
		w.Flush()
	},
}

func init() {
	fetchCmd.AddCommand(hintCmd)
}
