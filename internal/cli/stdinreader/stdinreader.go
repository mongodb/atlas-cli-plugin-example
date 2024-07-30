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

package stdinreader

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	return &cobra.Command{
		Use:   "stdinreader",
		Short: "Reads name and prints it",
		RunE: func(_ *cobra.Command, _ []string) error {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Please enter your name: ")
			name, err := reader.ReadString('\n')
			if err != nil {
				return err
			}

			name = strings.TrimSpace(name)
			fmt.Printf("Hello, %s!\n", name)
			return nil
		},
	}
}
