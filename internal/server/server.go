package server

import (
	"TaskProcessingService/internal/config"
	"TaskProcessingService/internal/handlers"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type IServer interface {
	Run(ctx context.Context) error
	setupRoutes()
}

type Server struct {
	router      *mux.Router
	server      *http.Server
	taskHandler handlers.ITaskHandler
	ctxlog      *log.Entry
}

func NewServer(tHandler handlers.ITaskHandler, config *config.ServerConfiguration) *Server {
	var (
		r *mux.Router
		s *http.Server
	)

	ctxlog := log.WithFields(log.Fields{
		"package":  "server",
		"function": "NewServer",
	})

	r = mux.NewRouter().StrictSlash(false)

	s = &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", config.HostPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return &Server{
		server:      s,
		router:      r,
		taskHandler: tHandler,
		ctxlog:      ctxlog,
	}

}

func (s *Server) Run(ctx context.Context) error {

	var err error
	ctxlog := s.ctxlog.WithFields(log.Fields{
		"function": "Run",
	})
	s.setupRoutes()

	serverErr := make(chan error, 1)

	go func() {
		ctxlog.Infof("HTTP server is running on %v", s.server.Addr)
		serverErr <- s.server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		ctxlog.Warn("gracefully shutdown the server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = s.server.Shutdown(ctx)
	case err = <-serverErr:
	}

	return fmt.Errorf("server error: %w", err)
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("/", s.taskHandler.ProcessTasks)
}
