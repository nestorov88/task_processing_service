package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	ServerConfig *ServerConfiguration
}

type ServerConfiguration struct {
	HostPort string `env:"SERVER_HOST_PORT,notEmpty"`
}

func NewConfiguration() (*Configuration, error) {
	ctxlog := log.WithFields(log.Fields{
		"package":  "config",
		"function": "NewConfiguration",
	})

	var (
		err          error
		serverConfig ServerConfiguration
	)

	if err = env.Parse(&serverConfig); err != nil {
		ctxlog.Warnf("could not parse server configuration: %s", err)

		return nil, fmt.Errorf("could not parse server configration: %w", err)
	}

	return &Configuration{ServerConfig: &serverConfig}, nil
}
