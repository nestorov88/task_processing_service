package models

import "golang.org/x/exp/slices"

type TasksRequest struct {
	Tasks *Tasks `json:"tasks"`
}

type TaskResponse struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type TasksResponse []TaskResponse

type Task struct {
	Name     string   `json:"name"`
	Command  string   `json:"command"`
	Requires []string `json:"requires"`
}

type Tasks []Task

func (t Tasks) Len() int {
	return len(t)
}
func (t Tasks) Less(i, j int) bool {
	return slices.Contains(t[j].Requires, t[i].Name)
}

func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
