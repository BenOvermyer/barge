package main

func main() {
	portainer := NewPortainer()

	portainer.token = portainer.login()

	endpoints := portainer.getEndpoints()
	endpoints = portainer.populateServicesForEndpoints(endpoints)
	endpoints = portainer.populateContainersForEndpoints(endpoints)
	endpoints = portainer.populateNetworksForEndpoints(endpoints)

	portainer.Endpoints = endpoints

	printEndpoints(portainer.Endpoints)
}
