package validation

import (
	"TaskProcessingService/internal/models"
	"testing"
)

func TestValidate(t *testing.T) {
	tValidator := &TaskValidator{}

	tests := []struct {
		name   string
		input  models.Tasks
		result error
	}{
		{
			name: "TestEmptyTaskNameValidationErr",
			input: models.Tasks{
				{
					Name:    "",
					Command: "echo 'Test'",
				},
			},
			result: ErrEmptyTaskName,
		},
		{
			name: "TestEmptyTaskCommandValidationErr",
			input: models.Tasks{
				{
					Name:    "task-1",
					Command: "",
				},
			},
			result: ErrEmptyTaskCommand,
		},
		{
			name: "TestNonExistingDependencyValidationErr",
			input: models.Tasks{
				{
					Name:    "task-1",
					Command: "test",
					Requires: []string{
						"task-2",
					},
				},
			},
			result: ErrNonExistingDependency,
		},
		{
			name: "TestCircularDependencyValidationErr",
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
					Requires: []string{
						"task-1",
					},
				},
			},
			result: ErrCircularDependency,
		},
		{
			name: "TestValidTasks",
			input: models.Tasks{
				{
					Name:    "task-1",
					Command: "test",
				},
				{
					Name:    "task-2",
					Command: "test",
					Requires: []string{
						"task-1",
					},
				},
			},
			result: ErrNonExistingDependency,
		},
		{
			name:   "TestEmptyTask",
			input:  models.Tasks{},
			result: ErrEmptyTasks,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := tValidator.Validate(test.input); err != nil {
				if err != test.result {
					t.Errorf("could not validate %v", err)
				}
			}
		})
	}
}
