package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

// Container is a Docker container
type Container struct {
	ID    string
	Image string
	State string
}

// Endpoint is a Docker Swarm endpoint
type Endpoint struct {
	ID         int
	Name       string
	Containers []Container
	Networks   []Network
	Services   []Service
}

// Network is a Docker network
type Network struct {
	Attachable bool
	Created    string
	ID         string
	Internal   bool
	Name       string
}

// Service is a Docker service
type Service struct {
	ID   string
	Spec struct {
		Name string
	}
}

type portainer struct {
	URL       string
	username  string
	password  string
	token     string
	endpoints []Endpoint
}

func (p portainer) fetch(item string) string {
	bearerHeader := "Bearer " + p.token
	requestURL := p.URL + "/" + item
	req, err := http.NewRequest("GET", requestURL, nil)

	req.Header.Set("Authorization", bearerHeader)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (p portainer) getContainersForEndpoint(endpoint Endpoint) []Container {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/containers/json")

	containers := make([]Container, 0)

	json.Unmarshal([]byte(output), &containers)

	return containers
}

func (p portainer) getEndpoints() []Endpoint {
	output := p.fetch("endpoints")

	endpoints := make([]Endpoint, 0)

	json.Unmarshal([]byte(output), &endpoints)

	return endpoints
}

func (p portainer) getNetworksForEndpoint(endpoint Endpoint) []Network {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/networks")

	networks := make([]Network, 0)

	json.Unmarshal([]byte(output), &networks)

	return networks
}

func printEndpoints(endpoints []Endpoint) {
	for _, e := range endpoints {
		fmt.Println(strconv.Itoa(e.ID) + ": " + e.Name + " (" + strconv.Itoa(len(e.Services)) + " services, " + strconv.Itoa(len(e.Containers)) + " containers, " + strconv.Itoa(len(e.Networks)) + " networks)")
	}
}

func printServices(services []Service) {
	for _, s := range services {
		fmt.Println(s.Spec.Name)
	}
}

func (p portainer) getServicesForEndpoint(endpoint Endpoint) []Service {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/services")

	services := make([]Service, 0)

	json.Unmarshal([]byte(output), &services)

	return services
}

func (p portainer) populateContainersForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Containers = p.getContainersForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func (p portainer) populateNetworksForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Networks = p.getNetworksForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func (p portainer) populateServicesForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Services = p.getServicesForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}

func (p portainer) login() string {
	var data map[string]interface{}

	token := ""
	url := p.URL + "/auth"

	authString := `{"Username": "` + p.username + `", "Password": "` + p.password + `"}`

	jsonBlock := []byte(authString)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBlock))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &data)

	token = data["jwt"].(string)

	return token
}

func main() {
	portainer := portainer{
		URL:      os.Getenv("PORTAINER_URL") + "/api",
		username: os.Getenv("PORTAINER_USERNAME"),
		password: os.Getenv("PORTAINER_PASSWORD"),
	}

	portainer.token = portainer.login()

	endpoints := portainer.getEndpoints()
	endpoints = portainer.populateServicesForEndpoints(endpoints)
	endpoints = portainer.populateContainersForEndpoints(endpoints)
	endpoints = portainer.populateNetworksForEndpoints(endpoints)

	portainer.endpoints = endpoints

	printEndpoints(portainer.endpoints)
}
