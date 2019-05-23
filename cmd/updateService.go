// Copyright © 2019 Ben Overmyer <ben@overmyer.net>
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
	"strconv"

	"github.com/spf13/cobra"
)

// updateServiceCmd represents the updateService command
var updateServiceCmd = &cobra.Command{
	Use:   "update",
	Short: "Trigger an update of the specified service",
	Long: `Trigger an update of the specified service.
	Note: this will update ALL services that match the given name/ID across all given endpoints.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("updateService called with " + args[0])

		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		path := ""

		data := map[string]interface{}{
			"TaskTemplate": map[string]interface{}{
				"ForceUpdate": 1,
			},
		}

		for _, e := range portainer.Endpoints {
			for _, s := range e.Services {
				if s.ID == args[0] || s.Spec.Name == args[0] {
					fmt.Println("Updating service " + s.Spec.Name + "...")
					path = "/services/" + s.ID + "/?version=" + strconv.Itoa(s.Version.Index)
					portainer.post(data, path)
				}
			}
		}
	},
}

func init() {
	serviceCmd.AddCommand(updateServiceCmd)
}
