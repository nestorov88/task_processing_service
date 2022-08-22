package json

import (
	"TaskProcessingService/internal/models"
	"encoding/json"
	"fmt"
	"io"
)

type TaskSerializer struct {
}

func (ts *TaskSerializer) Decode(r io.ReadCloser) (*models.TasksRequest, error) {
	var (
		err   error
		taskR *models.TasksRequest
	)

	if err = json.NewDecoder(r).Decode(&taskR); err != nil {
		err = fmt.Errorf("error while decoding request: %w", err)

		return nil, err
	}

	return taskR, nil
}

func (ts *TaskSerializer) Encode(input *models.TasksResponse) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("encoding error: %w", err)
	}
	return rawMsg, nil
}
