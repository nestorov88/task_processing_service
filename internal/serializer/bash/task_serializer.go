package bash

import (
	"TaskProcessingService/internal/models"
	"bytes"
	"errors"
	"fmt"
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
	b := bytes.NewBufferString("#!/usr/bin/env bash \n")

	for _, task := range *input {
		if _, err := b.WriteString(fmt.Sprintf("%s\n", task.Command)); err != nil {
			return nil, err
		}
	}

	return b.Bytes(), nil
}
