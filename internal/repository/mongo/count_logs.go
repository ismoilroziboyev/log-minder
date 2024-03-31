package mongo

import (
	"context"

	"github.com/ismoilroziboyev/log-minder/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongo) CountLogs(ctx context.Context, payload *domain.RetreiveLogsFilter) (int64, error) {
	return m.Database(m.cfg.MongoDB).Collection(collectionLogs).CountDocuments(ctx, bson.M(payload.Query))
}
