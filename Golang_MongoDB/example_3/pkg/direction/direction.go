package direction

import (
	"monGo/models"
//	"monGo/storage/mongoDB"
	"monGo/storage"
	"monGo/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HumanService struct {
	storage storage.Istorage
}

func NewHumanService (storage storage.Istorage) *HumanService {
	return &HumanService{
		storage: storage,
	}
}

func (h *HumanService) CreateHuman(human *models.Human) (*models.Human, *utils.RestErr) {
	human, restErr := h.storage.Human().Create(human)
	if restErr != nil {
		return nil, restErr
	}
	return human, nil
}

func (h *HumanService) FindHuman(id string) (*models.Human, *utils.RestErr) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
	  panic(err)
	}
	human, restErr := h.storage.Human().Find(objID)
	if restErr != nil {
		return nil, restErr
	}
	return human, nil
}

func (h *HumanService) DeleteHuman(id string) *utils.RestErr {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
	  panic(err)
	}
	restErr := h.storage.Human().Delete(objID)
	if restErr != nil {
		return restErr
	}
	return nil
}

func (h *HumanService) UpdateHuman(human *models.Human) (*models.Human, *utils.RestErr) {
	human, restErr := h.storage.Human().Update(human)
	if restErr != nil {
		return nil, restErr
	}

	return human, nil
}
