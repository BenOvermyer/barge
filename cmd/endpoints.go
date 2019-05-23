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

// Endpoint is a Docker Swarm endpoint
type Endpoint struct {
	ID         int
	Name       string
	Containers []Container
	Networks   []Network
	Services   []Service
	Tasks      []Task
	Nodes      []Node
}

func (p Portainer) getEndpoints() []Endpoint {
	endpoints := make([]Endpoint, 0)

	if endpointID != 0 {
		output := p.fetch("endpoints/" + strconv.Itoa(endpointID))

		endpoint := Endpoint{}

		json.Unmarshal([]byte(output), &endpoint)

		endpoints = append(endpoints, endpoint)
	} else {
		output := p.fetch("endpoints")

		json.Unmarshal([]byte(output), &endpoints)
	}

	endpoints = p.populateServicesForEndpoints(endpoints)
	endpoints = p.populateContainersForEndpoints(endpoints)
	endpoints = p.populateNetworksForEndpoints(endpoints)
	endpoints = p.populateNodesForEndpoints(endpoints)
	endpoints = p.populateTasksForEndpoints(endpoints)

	return endpoints
}

func (p Portainer) printEndpoints() {
	for _, e := range p.Endpoints {
		fmt.Println(strconv.Itoa(e.ID) + ": " + e.Name + " (" + strconv.Itoa(len(e.Services)) + " services, " + strconv.Itoa(len(e.Containers)) + " containers, " + strconv.Itoa(len(e.Networks)) + " networks)")
	}
	fmt.Println("----")
}
