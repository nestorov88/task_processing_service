package plain

import (
	"TaskProcessingService/internal/models"
	"bytes"
	"errors"
	"io"
)

var (
	ErrNotSupportedPlainTextDecoding = errors.New("plain text decoding not supported")
)

type TaskSerializer struct {
}

func (t TaskSerializer) Decode(r io.ReadCloser) (*models.TasksRequest, error) {
	return nil, ErrNotSupportedPlainTextDecoding
}

func (t TaskSerializer) Encode(input *models.TasksResponse) ([]byte, error) {
	var (
		b   bytes.Buffer
		err error
	)

	for _, task := range *input {
		if _, err = b.WriteString(task.Command + "\n"); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}
