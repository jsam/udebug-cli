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

// hintsCmd represents the hints command
var hintsCmd = &cobra.Command{
	Use:   "hints",
	Short: "Shows all available hints for given problem identifier",
	Long: `You must specify judge and problem identifiers as positional arguments. For example:

	udebug-cli show hints uva 495`,
	Run: func(cmd *cobra.Command, args []string) {
		judgeID, problemID := args[0], args[1]
		client := utils.NewAPIClient()


		params := hint.NewGetHintAPIHintListRetrieveJSONParams()
		params.JudgeAlias = judgeID
		params.ProblemID = problemID

		resp, err := client.Client.Hint.GetHintAPIHintListRetrieveJSON(params, client.Auth)
		if err != nil {
			fmt.Println(resp, err)
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', tabwriter.TabIndent)
		header := fmt.Sprintf("%s\t%s\t%s\t%s", "ID", "Votes", "User", "Date")
		fmt.Fprintln(w, header)

		for _, input := range resp.Payload {
			line := fmt.Sprintf("%s\t%s\t%s\t%s", input.ID, input.Votes, input.User, input.Date)
			fmt.Fprintln(w, line)
		}
		w.Flush()
	},
}

func init() {
	listCmd.AddCommand(hintsCmd)
}
