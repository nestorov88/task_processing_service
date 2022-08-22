package serializer

import (
	"TaskProcessingService/internal/models"
	"io"
)

type ITaskSerializer interface {
	Decode(r io.ReadCloser) (*models.TasksRequest, error)
	Encode(input *models.TasksResponse) ([]byte, error)
}
