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

package sharedlibrary

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20240805003/admin"
)

// Example modified from: https://github.com/mongodb/atlas-sdk-go/tree/main?tab=readme-ov-file#atlas-sdk-go
func makeAtlasCall() {
	ctx := context.Background()

	// This is the only difference from the example code
	// Use `config.HttpClient()` to get an authorized http client based on the default profile
	sdk, err := admin.NewClient(admin.UseHTTPClient(config.HttpClient()), admin.UseBaseURL(config.HttpBaseURL()))
	if err != nil {
		log.Fatalf("Error when instantiating new client: %v", err)
	}
	projects, response, err := sdk.ProjectsApi.ListProjects(ctx).Execute()
	if err != nil {
		log.Fatalf("Could not fetch projects: %v", err)
	}
	fmt.Printf("Response status: %v\n", response.Status)
	fmt.Printf("\nProject count: %v, Showing %v\n", *projects.TotalCount, len(*projects.Results))
	for _, project := range *projects.Results {
		fmt.Printf("- %s\n", project.Name)
	}
}

func CallAtlasBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "call-atlas",
		Short: "demo on how to call atlas.",
		Run: func(_ *cobra.Command, _ []string) {
			makeAtlasCall()
		},
	}

	return cmd
}
