package main

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
}

func (p Portainer) getEndpoints() []Endpoint {
	output := p.fetch("endpoints")

	endpoints := make([]Endpoint, 0)

	json.Unmarshal([]byte(output), &endpoints)

	return endpoints
}

func printEndpoints(endpoints []Endpoint) {
	for _, e := range endpoints {
		fmt.Println(strconv.Itoa(e.ID) + ": " + e.Name + " (" + strconv.Itoa(len(e.Services)) + " services, " + strconv.Itoa(len(e.Containers)) + " containers, " + strconv.Itoa(len(e.Networks)) + " networks)")
	}
}
