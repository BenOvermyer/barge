package main

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
