package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ismoilroziboyev/log-minder/internal/config"
	"github.com/ismoilroziboyev/log-minder/internal/repository"
	"github.com/ismoilroziboyev/log-minder/internal/services"
	"github.com/ismoilroziboyev/log-minder/internal/transport/http"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	cfg := config.Load()

	logrus.Info("environment variables read...")

	repo := repository.New(cfg)

	logrus.Info("repository initialized successfully...")

	service := services.New(repo, cfg)

	logrus.Info("service initalized sucessfully...")

	// initialize http server
	server := http.New(cfg, service)

	errChan := make(chan error, 1)

	go func() {
		logrus.Infof("LOG-MINDER application http-server started at: %s:%d....", cfg.HttpHost, cfg.HttpPort)

		if err := server.Run(); err != nil {
			errChan <- err
		}
	}()

	// implement gracefull shutdown
	q := make(chan os.Signal, 1)

	signal.Notify(q, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-q:
		logrus.Info("shutdown signal received from the system")

	case err := <-errChan:
		logrus.Infof("error received from starting up server, err: %s", err.Error())
	}

	logrus.Info("shutting down started...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.Infof("error occured while shutting down http server: %s", err.Error())
	}

	if err := repo.Close(ctx); err != nil {
		logrus.Infof("error occured while closing repository connections: %s", err.Error())
	}

	logrus.Info("server shutted down")

}
