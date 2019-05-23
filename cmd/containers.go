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

// Container is a Docker container
type Container struct {
	ID    string
	Image string
	State string
}

func (p Portainer) getContainersForEndpoint(endpoint Endpoint) []Container {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/containers/json")

	containers := make([]Container, 0)

	json.Unmarshal([]byte(output), &containers)

	return containers
}

func (p Portainer) populateContainersForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Containers = p.getContainersForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func printContainersForEndpoint(endpoint Endpoint) {
	fmt.Println("Containers in " + endpoint.Name)
	fmt.Println("----")

	for _, c := range endpoint.Containers {
		fmt.Println("ID: " + c.ID + ", image: " + c.Image)
	}
	fmt.Println("----")
}
