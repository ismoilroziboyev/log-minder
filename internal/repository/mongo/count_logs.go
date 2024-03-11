package mongo

import (
	"context"
	"fmt"
	"strings"

	"github.com/ismoilroziboyev/log-minder/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongo) CountLogs(ctx context.Context, payload *domain.RetreiveLogsFilter) (int64, error) {

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

	if payload.UserDetails != "" {
		list := strings.Split(payload.UserDetails, ",")

		for _, value := range list {
			query := strings.Split(value, "=")

			if len(query) != 2 {
				continue
			}

			filter[fmt.Sprintf("user.details.%s", query[0])] = query[1]
		}
	}

	if payload.ActionDetails != "" {
		list := strings.Split(payload.ActionDetails, ",")

		for _, value := range list {
			query := strings.Split(value, "=")

			if len(query) != 2 {
				continue
			}

			filter[fmt.Sprintf("action.details.%s", query[0])] = query[1]
		}
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

	return m.Database(m.cfg.MongoDB).Collection(collectionLogs).CountDocuments(ctx, filter)

}
