package services

import (
	"TaskProcessingService/internal/models"
	"TaskProcessingService/internal/sorter"
	"TaskProcessingService/internal/validation"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type ITaskService interface {
	GetSortedTasks(tReq *models.TasksRequest) (*models.TasksResponse, error)
}

type taskService struct {
	sorter    sorter.ITaskSorter
	validator validation.ITaskValidator
	ctxlog    *log.Entry
}

func NewTaskService(s sorter.ITaskSorter, v validation.ITaskValidator) *taskService {
	ctxlog := log.WithFields(log.Fields{
		"package":  "service",
		"function": "taskService",
	})

	ctxlog.Info("creating new instance of task service")
	return &taskService{
		sorter:    s,
		validator: v,
		ctxlog:    ctxlog,
	}
}

func (ts *taskService) GetSortedTasks(tReq *models.TasksRequest) (*models.TasksResponse, error) {
	ctxlog := ts.ctxlog.WithField("function", "GetSortedTasks")

	if err := ts.validator.Validate(*tReq.Tasks); err != nil {
		ctxlog.Errorf("validation error: %v", err)
		return nil, fmt.Errorf("validation error: %w", err)
	}

	if err := ts.sorter.Sort(*tReq.Tasks); err != nil {
		ctxlog.Errorf("sorting error: %v", err)

		return nil, fmt.Errorf("sorting error: %w", err)
	}

	taskResponse := make(models.TasksResponse, len(*tReq.Tasks))

	for _, task := range *tReq.Tasks {

		taskResponse = append(taskResponse, models.TaskResponse{
			Name:    task.Name,
			Command: task.Command,
		})
	}

	return &taskResponse, nil
}
