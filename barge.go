package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type portainer struct {
	URL string
	username string
	password string
	token string
}

func (p portainer) fetch(item string) string {
	req, err := http.NewRequest("GET", p.URL + "/" + item, nil)
	req.Header.Set("Authorization", "Bearer=" + p.token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func (p portainer) getEndpoints() string {
	output := p.fetch("endpoints")

	return output
}

func (p portainer) login() string {
	token := ""
	url := p.URL + "/auth"

	jsonBlock := []byte(`{"Username": "", "Password": ""}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBlock))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	token = string(body)

	return token
}

func main() {
	portainer := portainer{
		URL: os.Getenv("PORTAINER_URL"),
		username: os.Getenv("PORTAINER_USERNAME"),
		password: os.Getenv("PORTAINER_PASSWORD"),
	}

	portainer.token = portainer.login()

	fmt.Println(portainer.getEndpoints())
}