package v1

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"monGo/models"
	"monGo/pkg/direction"
	"monGo/pkg/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	direction direction.HumanService
}

func NewHumanLoad (direction direction.HumanService) *Handler {
	return &Handler{
		direction: direction,
	}
}

func (h *Handler) CreateHuman(c *gin.Context) {
	var newHuman models.Human
	if err := c.ShouldBindJSON(&newHuman); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}
	log.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	user, restErr := h.direction.CreateHuman(&newHuman)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *Handler) GetHuman(c *gin.Context) {
	log.Println("++++++++++++++++++++++++++++++++++++++++++++++++")
	humanId := c.Query("id")
	if humanId == "" {
		restErr := utils.BadRequest("no id")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, restErr := h.direction.FindHuman(humanId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateHuman(c *gin.Context) {
	humanId := c.Query("id")
	
	if humanId == "" {
		restErr := utils.BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	objID, err := primitive.ObjectIDFromHex(humanId)
	if err != nil {
	  panic(err)
	}
	var editedHuman models.Human
	editedHuman.ID = objID
	if err := c.ShouldBindJSON(&editedHuman); err != nil {
		restErr := utils.BadRequest("Invalid json.")
		c.JSON(restErr.Status, restErr)
		return
	}

	human, restErr := h.direction.UpdateHuman(&editedHuman)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, human)
}


func (h *Handler) DeleteHuman(c *gin.Context) {
	humanId := c.Query("id")
	if humanId == "" {
		restErr := utils.BadRequest("no email.")
		c.JSON(restErr.Status, restErr)
		return
	}
	restErr := h.direction.DeleteHuman(humanId)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, gin.H{"isRemoved": true})
}
