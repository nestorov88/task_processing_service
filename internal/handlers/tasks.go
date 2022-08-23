package handlers

import (
	"TaskProcessingService/internal/models"
	"TaskProcessingService/internal/serializer"
	"TaskProcessingService/internal/serializer/bash"
	jsn "TaskProcessingService/internal/serializer/json"
	"TaskProcessingService/internal/services"
	"encoding/json"
	log "github.com/sirupsen/logrus"

	"net/http"
)

type ITaskHandler interface {
	ProcessTasks(w http.ResponseWriter, r *http.Request)
}

type TaskHandler struct {
	taskService services.ITaskService
	ctxlog      *log.Entry
}

func NewTaskHandler(service services.ITaskService) *TaskHandler {
	ctxlog := log.WithFields(log.Fields{
		"package":  "handlers",
		"function": "NewTaskHandler",
	})

	ctxlog.Info("creating new instance of task handler")
	return &TaskHandler{taskService: service, ctxlog: ctxlog}
}

func (t *TaskHandler) ProcessTasks(w http.ResponseWriter, r *http.Request) {
	ctxlog := t.ctxlog.WithField("function", "ProcessTasks")

	contentType := r.Header.Get("Content-Type")
	acceptType := r.Header.Get("Accept")

	var (
		err   error
		taskR *models.TasksRequest
	)

	if taskR, err = getSerializer(contentType).Decode(r.Body); err != nil {
		ctxlog.Errorf("error while decoding request: %v", err)
		respondError(err, http.StatusBadRequest, w)

		return
	}

	tasks, err := t.taskService.GetSortedTasks(taskR)

	if err != nil {
		ctxlog.Errorf("GetSortedTasks error: : %v", err)
		respondError(err, http.StatusBadRequest, w)

		return
	}

	responseBody, err := getSerializer(acceptType).Encode(tasks)
	if err != nil {
		ctxlog.Error("error while writing response")
	}

	respond(w, acceptType, responseBody, http.StatusOK)
}

func respond(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	ctxlog := log.WithFields(log.Fields{
		"package":  "handlers",
		"function": "respond",
	})

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		ctxlog.Error("error while writing response")
	}
}

func respondError(err error, responseCode int, w http.ResponseWriter) {
	ctxlog := log.WithFields(log.Fields{
		"package":  "handlers",
		"function": "respondError",
	})
	w.Header().Add("Content-Type", "application/json")

	errResponse := struct {
		Message string
		Code    int
	}{
		Message: err.Error(),
		Code:    responseCode,
	}

	serializedResponse, err := json.Marshal(errResponse)

	if err != nil {
		ctxlog.Errorf("error while marshaling error response: %v", err)
	}
	if _, err = w.Write(serializedResponse); err != nil {
		ctxlog.Errorf("error while writing error response: %v", err)
	}
}

func getSerializer(contentType string) serializer.ITaskSerializer {
	if contentType == "text/plain" {
		return &bash.TaskSerializer{}
	}
	return &jsn.TaskSerializer{}
}
