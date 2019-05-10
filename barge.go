package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	endpointID := flag.Int("i", 0, "Endpoint ID to filter by. Defaults to no filtering (0).")
	showContainers := flag.Bool("c", false, "Print a list of containers.")
	showEndpoints := flag.Bool("e", false, "Print a list of endpoints.")
	showNetworks := flag.Bool("n", false, "Print a list of networks.")
	showServices := flag.Bool("s", false, "Print a list of services.")
	showBrokenServices := flag.Bool("b", false, "Print a list of broken services.")
	makePublic := flag.Bool("p", false, "Make all targeted resources public.")
	useVerboseMode := flag.Bool("v", false, "Use verbose mode.")

	flag.Parse()

	portainer := NewPortainer()
	portainer.token = portainer.login()
	if *useVerboseMode {
		portainer.verbose = true
	} else {
		portainer.verbose = false
	}

	if *endpointID != 0 {
		endpoint := portainer.getEndpoint(*endpointID)
		endpoint.Containers = portainer.getContainersForEndpoint(endpoint)
		endpoint.Networks = portainer.getNetworksForEndpoint(endpoint)
		endpoint.Services = portainer.getServicesForEndpoint(endpoint)
		endpoint.Tasks = portainer.getTasksForEndpoint(endpoint)

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

		if *showBrokenServices {
			printBrokenServicesForEndpoint(portainer.Endpoints[0])
		}
	} else {
		endpoints := portainer.getEndpoints()
		endpoints = portainer.populateServicesForEndpoints(endpoints)
		endpoints = portainer.populateContainersForEndpoints(endpoints)
		endpoints = portainer.populateNetworksForEndpoints(endpoints)
		endpoints = portainer.populateTasksForEndpoints(endpoints)

		portainer.Endpoints = endpoints

		if *showEndpoints {
			portainer.printEndpoints()
		}

		for _, e := range portainer.Endpoints {
			if *showContainers {
				printContainersForEndpoint(e)
			}

			if *showNetworks {
				printNetworksForEndpoint(e)
			}

			if *showServices {
				printServicesForEndpoint(e)
			}

			if *showBrokenServices {
				printBrokenServicesForEndpoint(e)
			}
		}
	}

	if *makePublic {
		cCount := 0
		cCountFailed := 0
		cCountTotal := 0
		nCount := 0
		nCountFailed := 0
		nCountTotal := 0
		sCount := 0
		sCountFailed := 0
		sCountTotal := 0
		for _, e := range portainer.Endpoints {
			for _, c := range e.Containers {
				if !strings.Contains(c.Image, "portainer") {
					if portainer.makePublic("container", c.ID) {
						cCount++
					} else {
						cCountFailed++
					}
				}
				cCountTotal++
			}
			for _, n := range e.Networks {
				if !strings.Contains(n.Name, "portainer") {
					if portainer.makePublic("network", n.ID) {
						nCount++
					} else {
						nCountFailed++
					}
				}
				nCountTotal++
			}
			for _, s := range e.Services {
				if !strings.Contains(s.Spec.Name, "portainer") {
					if portainer.makePublic("service", s.ID) {
						sCount++
					} else {
						sCountFailed++
					}
				}
				sCountTotal++
			}
		}
		fmt.Println("Made public: " + strconv.Itoa(cCount) + " of " + strconv.Itoa(cCountTotal) + " containers, " + strconv.Itoa(nCount) + " of " + strconv.Itoa(nCountTotal) + " networks, and " + strconv.Itoa(sCount) + " of " + strconv.Itoa(sCountTotal) + " services.")
	}
}
