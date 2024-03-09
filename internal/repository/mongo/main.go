package mongo

import (
	"context"

	"github.com/ismoilroziboyev/log-minder/internal/config"
	"github.com/ismoilroziboyev/log-minder/internal/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	driver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionLogs = "logs"
)

type MongoDB interface {
	InsertLog(ctx context.Context, payload *domain.InsertLogPayload) error
	Close(ctx context.Context) error
}

type mongo struct {
	*driver.Client
	cfg *config.Config
}

func New(cfg *config.Config) MongoDB {

	client, err := driver.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))

	if err != nil {
		logrus.Fatalf("cannot intialize mongo db client: %s", err.Error())
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		logrus.Fatalf("cannot send ping signal to mongodb: %s", err.Error())
	}

	logrus.Info("mongodb connected successfully...")

	keys := driver.IndexModel{
		Keys: bson.M{
			"created_at": -1, // 1 for ascending, -1 for descending
		},
	}

	// Create the index
	if _, err = client.Database(cfg.MongoDB).Collection(collectionLogs).Indexes().CreateOne(context.TODO(), keys); err != nil {
		logrus.Fatalf("cannot create index: %s", err.Error())
	}

	return &mongo{
		Client: client,
		cfg:    cfg,
	}
}

func (m *mongo) Close(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
