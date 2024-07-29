// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// func EchoBuilder() *cobra.Command {
// 	return &cobra.Command{
// 		Use: "echo",
// 		Short: "Echos the input args",
// 		RunE: func(_ *cobra.Command, args []string) error {
// 			if len(args) == 0 {
// 				return errors.New("no args given")
// 			}
// 			for _, arg := range args {
// 				fmt.Printf("%s ", arg)
// 			}
// 			return nil
// 		},
// 	}
// }

// func HelloBuilder() *cobra.Command {
// 	return &cobra.Command{
// 		Use: "hello",
// 		Short: "The Hello World command",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			fmt.Println("Hello World!")
// 		},
// 	}
// }

func main() {
	var rootCmd = &cobra.Command{
		Use:   "example",
		Short: "Root command of the atlas cli plugin example",
	}

	// rootCmd.AddCommand(
	// 	HelloBuilder(),
	// 	EchoBuilder(),
	// )

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}