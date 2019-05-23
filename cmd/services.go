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

// Service is a Docker service
type Service struct {
	ID      string
	Version struct {
		Index int
	}
	Spec struct {
		Name string
		Mode struct {
			Replicated struct {
				Replicas int
			}
			Global string
		}
		Labels       map[string]string
		TaskTemplate struct {
			ContainerSpec struct {
				Env []string
			}
		}
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

func (e Endpoint) getBrokenServices() []Service {
	services := []Service{}

	for _, s := range e.Services {
		if e.getServiceTaskStatus(s) == "broken" {
			services = append(services, s)
		}
	}

	return services
}

func printBrokenServicesForEndpoint(endpoint Endpoint) {
	brokenServices := endpoint.getBrokenServices()

	fmt.Println("Broken services for " + endpoint.Name)
	fmt.Println("----")

	for _, s := range brokenServices {
		fmt.Println(s.Spec.Name + " (" + endpoint.getReplicaStatusForService(s) + ")")
	}
}

func printServicesForEndpoint(endpoint Endpoint) {
	fmt.Println("Services in " + endpoint.Name)
	fmt.Println("----")

	for _, s := range endpoint.Services {
		fmt.Println("Name: " + s.Spec.Name + ", ID: " + s.ID)
	}
	fmt.Println("----")
}

func printServiceLabelsForEndpoint(endpoint Endpoint) {
	fmt.Println("Service Labels in " + endpoint.Name)
	fmt.Println("----")

	for _, s := range endpoint.Services {
		fmt.Println("+-- Service Name: " + s.Spec.Name + ", ID: " + s.ID)
		for k, l := range s.Spec.Labels {
			fmt.Println("   Label: " + k + "=" + l)
		}
	}
	fmt.Println("----")
}

func printServiceVariablesForEndpoint(endpoint Endpoint) {
	fmt.Println("Service Enviroment Variables in " + endpoint.Name)
	fmt.Println("----")

	for _, s := range endpoint.Services {
		fmt.Println("+-- Service Name: " + s.Spec.Name + ", ID: " + s.ID)
		for _, ev := range s.Spec.TaskTemplate.ContainerSpec.Env {
			fmt.Println("   Variable: " + ev)
		}
	}
	fmt.Println("----")
}
