package services

import (
	"TaskProcessingService/internal/mocks"
	"TaskProcessingService/internal/models"
	"TaskProcessingService/internal/validation"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetSortedTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	srt := mocks.NewMockITaskSorter(ctrl)
	validator := mocks.NewMockITaskValidator(ctrl)
	service := NewTaskService(srt, validator)

	tests := []struct {
		name            string
		input           *models.TasksRequest
		sorterResult    error
		validatorResult error
	}{
		{
			name: "GetSortedTaskNoError",
			input: &models.TasksRequest{
				Tasks: &models.Tasks{
					models.Task{
						Name:    "task-test",
						Command: "echo 'test'",
					},
				},
			},
		},
		{
			name: "GetSortedTaskSorterError",
			input: &models.TasksRequest{
				Tasks: &models.Tasks{},
			},
			sorterResult: errors.New("sorting error"),
		},
		{
			name: "GetSortedTaskValidationError",
			input: &models.TasksRequest{
				Tasks: &models.Tasks{},
			},
			validatorResult: validation.ErrEmptyTaskName,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if test.validatorResult == nil {
				srt.EXPECT().Sort(*test.input.Tasks).Return(test.sorterResult)
			}

			validator.EXPECT().Validate(*test.input.Tasks).Return(test.validatorResult)

			if _, err := service.GetSortedTasks(test.input); err != nil && (test.sorterResult == nil && test.validatorResult == nil) {
				t.Errorf("error getting sorted tasks: %v", err)
			}
		})
	}

}
