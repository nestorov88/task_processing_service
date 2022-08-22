package app

import (
	"TaskProcessingService/internal/config"
	"TaskProcessingService/internal/handlers"
	srv "TaskProcessingService/internal/server"
	"TaskProcessingService/internal/services"
	"TaskProcessingService/internal/sorter"
	"TaskProcessingService/internal/validation"
	"context"
	log "github.com/sirupsen/logrus"
	"sync"
)

type App struct {
	server srv.IServer
	ctxlog *log.Entry
}

func NewApp(c *config.Configuration) *App {
	ctxlog := log.WithFields(log.Fields{
		"package":  "app",
		"function": "NewApp",
	})

	ctxlog.Info("creating new instance of app")

	var (
		service services.ITaskService
		handler handlers.ITaskHandler
		server  srv.IServer
	)

	service = services.NewTaskService(sorter.DependencyTaskSorter{}, validation.TaskValidator{})

	handler = handlers.NewTaskHandler(service)

	server = srv.NewServer(handler, c.ServerConfig)

	return &App{server: server, ctxlog: ctxlog}
}

func (a *App) Run(ctx context.Context) error {
	var (
		err error
		wg  sync.WaitGroup
	)

	ctxlog := a.ctxlog.WithField("function", "Run")
	ctxlog.Info("app starts running")

	wg.Add(1)

	go func() {
		err = a.server.Run(ctx)
		wg.Done()
	}()

	wg.Wait()

	return err
}
