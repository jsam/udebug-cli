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
	"os/exec"
	"bufio"
	"bytes"
	"strings"
	"github.com/sergi/go-diff/diffmatchpatch"
	"fmt"
	"os"
)

// solutionCmd represents the solution command
var solutionCmd = &cobra.Command{
	Use:   "solution",
	Short: "Executes specified solution and checks if output is correct.",
	Long: `You must specify path to your solution and input identifier as positional arguments. For example:

	udebug-cli test solution $HOME/solutions/fibb/fibbonaci.bin 809839`,
	Run: func(cmd *cobra.Command, args []string) {
		solutionPath, inputID := args[0], args[1]
		inputSlice := getInput(inputID)
		inputStr := []byte(strings.Join(inputSlice, ""))
		correctOutputStr := getOutput(inputID)

		handlerScript := exec.Command(solutionPath)
		handlerScript.Stdin =  bufio.NewReader(bytes.NewBuffer(inputStr))
		output, err := handlerScript.Output()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(string(output), correctOutputStr, true)
		report := dmp.DiffPrettyText(diffs)
		fmt.Println(report)

		if diffs[0].Type == diffmatchpatch.DiffEqual {
			fmt.Println("  🎉   Your solution is correct.")
		} else {
			fmt.Println("  😔   Your solution is wrong.")
		}
	},
}

func init() {
	testCmd.AddCommand(solutionCmd)
}
