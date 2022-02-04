package db

import (
	"context"
	"time"

	"monGo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnDB(cfg config.Config) (*mongo.Database, error) {
	uri := cfg.MongoURI
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := client.Database(cfg.MongoDatabase)

	return db, nil
}
