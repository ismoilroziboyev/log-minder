package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ismoilroziboyev/log-minder/internal/config"
	"github.com/ismoilroziboyev/log-minder/internal/services"
	"github.com/ismoilroziboyev/log-minder/internal/transport/http/handlers"
)

type HttpServer struct {
	*http.Server
	cfg *config.Config
}

func New(cfg *config.Config, services *services.Service) *HttpServer {
	// initializa handlers
	handlers := handlers.New(cfg, services)

	return &HttpServer{
		Server: &http.Server{
			Addr:           fmt.Sprintf("%s:%d", cfg.HttpHost, cfg.HttpPort),
			ReadTimeout:    time.Second * 15,
			WriteTimeout:   time.Second * 15,
			MaxHeaderBytes: 1 << 20,
			Handler:        handlers.API(),
		},
		cfg: cfg,
	}
}

// @title LOG-MINDER APPLICATION API DOCUMENTATION
// @description This API contains the source for the LOG-MINDER  app

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

// @BasePath /api/v1

// Run initializes http server
func (s *HttpServer) Run() error {
	return s.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
