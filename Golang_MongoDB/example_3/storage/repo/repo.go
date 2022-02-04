package repo

import (
	"monGo/models"
	"monGo/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HumanStorageI...
type HumanStorageI interface {
	Create(*models.Human) (*models.Human, *utils.RestErr)
	Find(id primitive.ObjectID) (*models.Human, *utils.RestErr)
	Update(*models.Human) (*models.Human, *utils.RestErr)
	Delete(id primitive.ObjectID) *utils.RestErr
}