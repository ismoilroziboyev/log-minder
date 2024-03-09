package mongo

import (
	"context"
	"time"

	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
	"github.com/sirupsen/logrus"
)

func (m *mongo) InsertLog(ctx context.Context, payload *domain.InsertLogPayload) error {

	payload.CreatedAt = time.Now()

	res, err := m.Database(m.cfg.MongoDB).Collection(collectionLogs).InsertOne(ctx, payload)

	if err != nil {
		return errors.NewInternalServerErrorw("cannot save log", err)
	}

	logrus.Infof("log saved with id: %v", res.InsertedID)

	return nil
}
