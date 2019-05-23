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
	"github.com/spf13/cobra"
)

var filterServicesBroken bool

// listServicesCmd represents the listServices command
var listServicesCmd = &cobra.Command{
	Use:   "list",
	Short: "List services",
	Long:  `List all services running in the specified endpoint(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		portainer := NewPortainer()
		portainer.Endpoints = portainer.getEndpoints()

		for _, e := range portainer.Endpoints {
			if filterServicesBroken {
				printBrokenServicesForEndpoint(e)
			} else {
				printServicesForEndpoint(e)
			}
		}
	},
}

func init() {
	serviceCmd.AddCommand(listServicesCmd)
	listServicesCmd.Flags().BoolVarP(&filterServicesBroken, "broken", "b", false, "Display only broken services")
}
