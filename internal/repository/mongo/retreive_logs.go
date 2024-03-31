package mongo

import (
	"context"

	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongo) RetreiveLogs(ctx context.Context, payload *domain.RetreiveLogsFilter) ([]domain.Log, error) {

	if payload.Limit <= 0 {
		payload.Limit = 15
	}

	if payload.Offset < 0 {
		payload.Offset = 0
	}
	cur, err := m.Database(m.cfg.MongoDB).
		Collection(collectionLogs).
		Find(ctx, bson.M(payload.Query), options.
			Find().
			SetSort(map[string]int{"created_at": -1}).
			SetSkip(int64(payload.Offset)).
			SetLimit(int64(payload.Limit)))

	if err != nil {
		return nil, errors.NewInternalServerErrorw("cannot get documents", err)
	}

	var logs []domain.Log

	if err := cur.All(ctx, &logs); err != nil {
		return nil, errors.NewInternalServerErrorw("cannot decode documents to results", err)
	}

	//close the cursor
	cur.Close(ctx)

	return logs, nil
}
