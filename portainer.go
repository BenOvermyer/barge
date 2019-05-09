package main

import "os"

// Portainer is an instance of Portainer
type Portainer struct {
	URL       string
	username  string
	password  string
	token     string
	Endpoints []Endpoint
}

// NewPortainer returns a new Portainer instance from the environment
func NewPortainer() Portainer {
	portainer := Portainer{
		URL:      os.Getenv("PORTAINER_URL") + "/api",
		username: os.Getenv("PORTAINER_USERNAME"),
		password: os.Getenv("PORTAINER_PASSWORD"),
	}

	return portainer
}
