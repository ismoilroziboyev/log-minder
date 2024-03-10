package logs

import (
	"context"
	"time"

	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
)

func (s *service) Retreive(ctx context.Context, payload *domain.RetreiveLogsFilter) (*domain.RetreiveLogsResponse, error) {
	var timeFormat = "2006-01-02"

	if payload.From != "" {

		from, err := time.Parse(timeFormat, payload.From)

		if err != nil {
			return nil, errors.NewBadRequestErrorf("invalid from time format, time layout is: %s", timeFormat)
		}
		payload.FromDate = from
	}

	if payload.To != "" {
		to, err := time.Parse(timeFormat, payload.To)

		if err != nil {
			return nil, errors.NewBadRequestErrorf("invalid to time format, time layout is: %s", timeFormat)
		}

		payload.ToDate = to
	}

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
