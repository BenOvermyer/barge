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
	"encoding/json"
	"fmt"
	"strconv"
)

// Network is a Docker network
type Network struct {
	Attachable bool
	Created    string
	ID         string
	Internal   bool
	Name       string
}

func (p Portainer) getNetworksForEndpoint(endpoint Endpoint) []Network {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/networks")

	networks := make([]Network, 0)

	json.Unmarshal([]byte(output), &networks)

	return networks
}

func (p Portainer) populateNetworksForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Networks = p.getNetworksForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func printNetworksForEndpoint(endpoint Endpoint) {
	fmt.Println("Networks in " + endpoint.Name)
	fmt.Println("----")

	for _, n := range endpoint.Networks {
		fmt.Println("Name: " + n.Name + ", ID: " + n.ID)
	}
	fmt.Println("----")
}
