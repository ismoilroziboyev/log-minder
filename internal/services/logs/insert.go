package logs

import (
	"context"

	"github.com/ismoilroziboyev/log-minder/internal/domain"
)

func (s *service) Insert(ctx context.Context, payload *domain.InsertLogPayload) error {
	return s.repo.MongoDB.InsertLog(ctx, payload)
}
