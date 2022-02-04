package storage

import (
	"go.mongodb.org/mongo-driver/mongo"

	"monGo/storage/mongoDB"
	"monGo/storage/repo"
)

// istorage
type Istorage interface {
	Human() repo.HumanStorageI
}

type storagePg struct {
	db        *mongo.Database
	orderRepo repo.HumanStorageI
}

func NewStorageMg(db *mongo.Database) *storagePg {
	return &storagePg{
		db:        db,
		orderRepo: mongoDB.NewHumanRepo(db),
	}
}

func (s storagePg) Human() repo.HumanStorageI {
	return s.orderRepo
}