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
	"fmt"
	"os"

	"atlas-cli-plugin/internal/cli/echo"
	"atlas-cli-plugin/internal/cli/hello"
	"atlas-cli-plugin/internal/cli/printenv"
	"atlas-cli-plugin/internal/cli/sharedlibrary"
	"atlas-cli-plugin/internal/cli/stdinreader"

	"github.com/spf13/cobra"
)

func main() {
	exampleCmd := &cobra.Command{
		Use:     "example",
		Short:   "Root command of the atlas cli plugin example",
		Aliases: []string{"example-alias-1", "example-alias-2"},
	}

	exampleCmd.AddCommand(
		hello.Builder(),
		echo.Builder(),
		printenv.Builder(),
		stdinreader.Builder(),
		sharedlibrary.Builder(),
	)

	completionOption := &cobra.CompletionOptions{
		DisableDefaultCmd:   true,
		DisableNoDescFlag:   true,
		DisableDescriptions: true,
		HiddenDefaultCmd:    true,
	}
	rootCmd := &cobra.Command{
		DisableFlagParsing: true,
		DisableAutoGenTag:  true,
		DisableSuggestions: true,
		CompletionOptions:  *completionOption,
	}
	rootCmd.AddCommand(exampleCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
