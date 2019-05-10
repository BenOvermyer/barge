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
	Tasks      []Task
}

func (p Portainer) getEndpoint(id int) Endpoint {
	output := p.fetch("endpoints/" + strconv.Itoa(id))

	endpoint := Endpoint{}

	json.Unmarshal([]byte(output), &endpoint)

	return endpoint
}

func (p Portainer) getEndpoints() []Endpoint {
	output := p.fetch("endpoints")

	endpoints := make([]Endpoint, 0)

	json.Unmarshal([]byte(output), &endpoints)

	return endpoints
}

func (p Portainer) printEndpoints() {
	for _, e := range p.Endpoints {
		fmt.Println(strconv.Itoa(e.ID) + ": " + e.Name + " (" + strconv.Itoa(len(e.Services)) + " services, " + strconv.Itoa(len(e.Containers)) + " containers, " + strconv.Itoa(len(e.Networks)) + " networks)")
	}
	fmt.Println("----")
}
