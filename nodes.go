package main

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
