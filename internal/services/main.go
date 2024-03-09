package services

import (
	"github.com/ismoilroziboyev/log-minder/internal/config"
	"github.com/ismoilroziboyev/log-minder/internal/repository"
	"github.com/ismoilroziboyev/log-minder/internal/services/logs"
)

type Service struct {
	logs.LogsService
}

func New(repo *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		LogsService: logs.New(repo, cfg),
	}
}
