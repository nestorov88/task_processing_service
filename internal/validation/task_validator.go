package validation

import (
	"TaskProcessingService/internal/models"
	"errors"
	"golang.org/x/exp/slices"
)

type ITaskValidator interface {
	Validate(tasks models.Tasks) error
}

var (
	ErrEmptyTaskName         = errors.New("empty task name")
	ErrEmptyTaskCommand      = errors.New("empty task command ")
	ErrCircularDependency    = errors.New("task circular dependency")
	ErrNonExistingDependency = errors.New("task non existing dependency")
	ErrEmptyTasks            = errors.New("empty tasks list")
)

type TaskValidator struct {
}

func (tv TaskValidator) Validate(tasks models.Tasks) error {

	if len(tasks) == 0 {
		return ErrEmptyTasks
	}

	mTasks := make(map[string][]string, len(tasks))

	for _, task := range tasks {
		if task.Name == "" {
			return ErrEmptyTaskName
		}

		if task.Command == "" {
			return ErrEmptyTaskCommand
		}

		mTasks[task.Name] = task.Requires
	}

	for _, task := range tasks {

		for _, depTask := range task.Requires {

			dep, ok := mTasks[depTask]
			if ok != true {
				return ErrNonExistingDependency
			}

			if slices.Contains(dep, task.Name) {
				return ErrCircularDependency
			}

		}
	}

	return nil
}
