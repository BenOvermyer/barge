package main

import (
	"encoding/json"
	"strconv"
)

// Task is a Docker task
type Task struct {
	Name      string
	ServiceID string
	NodeID    string
	Status    struct {
		State   string
		Message string
	}
	DesiredState string
}

func (p Portainer) getTasksForEndpoint(endpoint Endpoint) []Task {
	output := p.fetch("endpoints/" + strconv.Itoa(endpoint.ID) + "/docker/tasks")

	tasks := make([]Task, 0)

	json.Unmarshal([]byte(output), &tasks)

	return tasks
}

func (e Endpoint) findTasksForService(service Service) []Task {
	tasks := []Task{}

	for _, t := range e.Tasks {
		if t.ServiceID == service.ID {
			tasks = append(tasks, t)
		}
	}

	return tasks
}

func (e Endpoint) getServiceTaskStatus(service Service) string {
	tasks := e.findTasksForService(service)

	running := 0
	desired := service.Spec.Mode.Replicated.Replicas
	global := service.Spec.Mode.Global

	for _, t := range tasks {
		if t.Status.State == "running" && t.DesiredState == "running" {
			running++
		}
	}

	if desired > 0 {
		if running == desired {
			return "working"
		}
		return "broken"
	}

	return global
}

func (e Endpoint) getReplicaStatusForService(service Service) string {
	tasks := e.findTasksForService(service)

	running := 0
	desired := service.Spec.Mode.Replicated.Replicas

	for _, t := range tasks {
		if t.Status.State == "running" && t.DesiredState == "running" {
			running++
		}
	}

	return strconv.Itoa(running) + "/" + strconv.Itoa(desired)
}

func (p Portainer) populateTasksForEndpoints(endpoints []Endpoint) []Endpoint {
	newEndpoints := []Endpoint{}
	var endpoint Endpoint

	for _, e := range endpoints {
		endpoint = e
		endpoint.Tasks = p.getTasksForEndpoint(e)

		newEndpoints = append(newEndpoints, endpoint)
	}

	return newEndpoints
}
