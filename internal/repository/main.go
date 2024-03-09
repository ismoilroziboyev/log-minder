package repository

import (
	"context"

	"github.com/ismoilroziboyev/log-minder/internal/config"
	"github.com/ismoilroziboyev/log-minder/internal/repository/mongo"
)

type Repository struct {
	mongo.MongoDB
}

func New(cfg *config.Config) *Repository {
	return &Repository{
		MongoDB: mongo.New(cfg),
	}
}

func (r *Repository) Close(ctx context.Context) error {

	if err := r.MongoDB.Close(ctx); err != nil {
		return err
	}

	return nil
}
