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

// Node is a Docker Swarm node (manager or worker)
type Node struct {
	ID   string
	Spec struct {
		Role         string
		Availability string
	}
	Description struct {
		Hostname string
	}
	Status struct {
		State string
		Addr  string
	}
	ManagerStatus struct {
		Reachability string
		Addr         string
	}
}

func (p Portainer) getNodesForEndpoint(endpoint Endpoint) []Node {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/nodes")

	nodes := make([]Node, 0)

	json.Unmarshal([]byte(output), &nodes)

	return nodes
}

func (p Portainer) populateNodesForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Nodes = p.getNodesForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func printNodesForEndpoint(endpoint Endpoint) {
	fmt.Println("Nodes in " + endpoint.Name)
	fmt.Println("----")

	for _, n := range endpoint.Nodes {
		fmt.Println("(" + n.Spec.Role + ") Hostname: " + n.Description.Hostname + ", IP: " + n.Status.Addr + " - this node is " + n.Status.State)
	}
	fmt.Println("----")
}
