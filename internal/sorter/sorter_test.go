package sorter

import (
	"TaskProcessingService/internal/models"
	"testing"
)

func TestSort(t *testing.T) {

	sorter := &DependencyTaskSorter{}

	tests := []struct {
		name   string
		input  models.Tasks
		result models.Tasks
	}{
		{
			input: models.Tasks{
				{
					Name:    "task-1",
					Command: "test",
					Requires: []string{
						"task-2",
					},
				},
				{
					Name:    "task-2",
					Command: "test",
				},
			},
			result: models.Tasks{
				{
					Name:    "task-2",
					Command: "test",
				},
				{
					Name:    "task-1",
					Command: "test",
					Requires: []string{
						"task-2",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := sorter.Sort(test.input); err != nil {
				t.Errorf("error while sorting %v", err)
			}
			for i, sortedTask := range test.input {
				if sortedTask.Name != test.result[i].Name {
					t.Errorf("wrong sorting")
				}
			}
		})
	}
}
