package api

import (

	"log"
	"github.com/gin-gonic/gin"

	m "monGo/middlewares"
	"monGo/api/handlers/v1"
)

func InitRoutes (r *gin.Engine, h *v1.Handler) {

	r.Use(m.Cors())

	// CATEGORIES
	log.Println("++++++++++++++++++++++++++++++++")
	router := r.Group("/v1")

	log.Println("++++++++++++++++++++++++++++++++")

	router.POST("/human/create", h.CreateHuman)
	router.GET("/human/find", h.GetHuman)
	router.PUT("/human/update", h.UpdateHuman)
	router.DELETE("/human/delete", h.DeleteHuman)
	
}
