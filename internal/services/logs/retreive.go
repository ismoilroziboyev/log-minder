package logs

import (
	"context"

	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
)

func (s *service) Retreive(ctx context.Context, payload *domain.RetreiveLogsFilter) (*domain.RetreiveLogsResponse, error) {
	count, err := s.repo.MongoDB.CountLogs(ctx, payload)

	if err != nil {
		return nil, errors.NewInternalServerErrorw("cannot get count of logs", err)
	}

	logs, err := s.repo.MongoDB.RetreiveLogs(ctx, payload)

	if err != nil {
		return nil, errors.NewInternalServerErrorw("cannot get logs", err)
	}

	var resp domain.RetreiveLogsResponse

	resp.Logs = logs

	resp.SetPaginationMetadata(payload.Limit, payload.Offset, count)

	return &resp, nil
}
