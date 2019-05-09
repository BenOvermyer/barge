package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Service is a Docker service
type Service struct {
	ID   string
	Spec struct {
		Name string
	}
}

func printServices(services []Service) {
	for _, s := range services {
		fmt.Println(s.Spec.Name)
	}
}

func (p Portainer) getServicesForEndpoint(endpoint Endpoint) []Service {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/services")

	services := make([]Service, 0)

	json.Unmarshal([]byte(output), &services)

	return services
}

func (p Portainer) populateServicesForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Services = p.getServicesForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func printServicesForEndpoint(endpoint Endpoint) {
	fmt.Println("Services in " + endpoint.Name)
	fmt.Println("----")

	for _, s := range endpoint.Services {
		fmt.Println("Name: " + s.Spec.Name + ", ID: " + s.ID)
	}
	fmt.Println("----")
}
