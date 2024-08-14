// Copyright 2020 MongoDB Inc
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

package listprofiles

import (
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/spf13/cobra"
)


func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "listprofiles",
		Short:   "Return a list of available profiles by name.",
		Long:    `If you did not specify a name for your profile, it displays as the default profile.`,
		Example: "  atlas config ls",
		Run: func(_ *cobra.Command, _ []string) {
			profiles := config.List()

			for _, profile := range profiles {
				fmt.Printf("%s\n", profile)
			}
		},
	}

	return cmd
}
