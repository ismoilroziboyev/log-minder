package logs

import (
	"context"

	"github.com/ismoilroziboyev/log-minder/internal/config"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
	"github.com/ismoilroziboyev/log-minder/internal/repository"
)

type LogsService interface {
	Insert(ctx context.Context, payload *domain.InsertLogPayload) error
	Retreive(ctx context.Context, payload *domain.RetreiveLogsFilter) (*domain.RetreiveLogsResponse, error)
}

type service struct {
	repo *repository.Repository
	cfg  *config.Config
}

func New(repo *repository.Repository, cfg *config.Config) LogsService {
	return &service{
		repo: repo,
		cfg:  cfg,
	}
}
