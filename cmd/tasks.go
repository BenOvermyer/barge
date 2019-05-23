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
