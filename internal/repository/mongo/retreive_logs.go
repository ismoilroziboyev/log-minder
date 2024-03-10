package mongo

import (
	"context"

	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *mongo) RetreiveLogs(ctx context.Context, payload *domain.RetreiveLogsFilter) ([]domain.Log, error) {

	var filter = bson.M{}

	if payload.ActionType != "" {
		filter["action.type"] = payload.ActionType
	}

	if payload.UserID != "" {
		filter["user.user_id"] = payload.UserID
	}

	if payload.UserRole != "" {
		filter["user.role"] = payload.UserRole
	}

	if payload.Search != "" {
		filter["$or"] = []bson.M{
			{"user.full_name": bson.M{"$regex": payload.Search, "$options": "i"}},
		}
	}

	if !payload.FromDate.IsZero() {
		filter["created_at"] = bson.M{"$gt": payload.FromDate}
	}

	if !payload.ToDate.IsZero() {
		filter["created_at"] = bson.M{"$lt": payload.ToDate}
	}

	cur, err := m.Database(m.cfg.MongoDB).Collection(collectionLogs).Find(ctx, filter, options.Find().SetSkip(int64(payload.Offset)).SetLimit(int64(payload.Limit)))

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
