package main

import (
	"TaskProcessingService/internal/app"
	"TaskProcessingService/internal/config"
	"context"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	// Output to stdout instead of the default stderr.
	log.SetOutput(os.Stdout)

	var (
		err error
		a   *app.App
		c   *config.Configuration
	)

	ctxlog := log.WithFields(log.Fields{
		"package":  "main",
		"function": "main",
	})

	_ = godotenv.Load()

	if c, err = config.NewConfiguration(); err != nil {
		ctxlog.Panicf("could not create new configuration: %v", err)
	}

	if a, err = app.NewApp(c); err != nil {
		ctxlog.Panicf("could not create new application: %v", err)
	}

	if err = a.Run(context.Background()); err != nil {
		ctxlog.Panicf("app stopped running: %v", err)
	} 
}
