package main

import (
	"flag"
	"fmt"
)

func main() {
	endpointID := flag.Int("eid", 0, "Endpoint ID to filter by")
	showContainers := flag.Bool("containers", false, "Print a list of containers")
	showEndpoints := flag.Bool("endpoints", false, "Print a list of endpoints")
	showNetworks := flag.Bool("networks", false, "Print a list of networks")
	showServices := flag.Bool("services", false, "Print a list of services")

	flag.Parse()

	portainer := NewPortainer()
	portainer.token = portainer.login()

	if *endpointID != 0 {
		endpoint := portainer.getEndpoint(*endpointID)
		endpoint.Containers = portainer.getContainersForEndpoint(endpoint)
		endpoint.Networks = portainer.getNetworksForEndpoint(endpoint)
		endpoint.Services = portainer.getServicesForEndpoint(endpoint)

		endpoints := []Endpoint{}
		endpoints = append(endpoints, endpoint)
		portainer.Endpoints = endpoints

		if *showEndpoints {
			portainer.printEndpoints()
		}

		if *showContainers {
			printContainersForEndpoint(portainer.Endpoints[0])
		}

		if *showNetworks {
			printNetworksForEndpoint(portainer.Endpoints[0])
		}

		if *showServices {
			printServicesForEndpoint(portainer.Endpoints[0])
		}
	} else {
		endpoints := portainer.getEndpoints()
		endpoints = portainer.populateServicesForEndpoints(endpoints)
		endpoints = portainer.populateContainersForEndpoints(endpoints)
		endpoints = portainer.populateNetworksForEndpoints(endpoints)

		portainer.Endpoints = endpoints

		if *showEndpoints {
			portainer.printEndpoints()
		} else {
			fmt.Println("Not doing much here...")
		}
	}
}
