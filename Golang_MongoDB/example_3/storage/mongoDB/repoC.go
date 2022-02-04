package mongoDB

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type humanRepo struct {
	db *mongo.Database
}

// NewOrderRepo... .
func NewHumanRepo(db *mongo.Database) *humanRepo {
	return &humanRepo{db: db}
}
