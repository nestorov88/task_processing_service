package sorter

import (
	"TaskProcessingService/internal/models"
	"golang.org/x/exp/slices"
)

type ITaskSorter interface {
	Sort(tasks models.Tasks) error
}

type DependencyTaskSorter struct {
}

func (d DependencyTaskSorter) Sort(tasks models.Tasks) error {

	for i := 0; i < len(tasks)-1; i++ {
		for j := 0; j < len(tasks)-i-1; j++ {
			if slices.Contains(tasks[j].Requires, tasks[j+1].Name) {
				tasks[j], tasks[j+1] = tasks[j+1], tasks[j]
			}

		}
	}

	return nil
}
