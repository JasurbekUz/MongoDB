package main

import (
	"log"

	"monGo/api"
	"monGo/config"
	"monGo/pkg/db"
	"monGo/pkg/direction"
	"monGo/api/handlers/v1"
	"monGo/storage"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main() {
	// load config
	cfg := config.Load()

	// connect mongodb
	connDB, err := db.ConnDB(cfg)
	if err != nil {
		log.Fatal("mongo connection to mongodb error")
	}
	// load storage
	mgStorage := storage.NewStorageMg(connDB)

	humanService := direction.NewHumanService(*mgStorage)

	hmHandler := v1.NewHumanLoad(*humanService)

	api.InitRoutes(router, *hmHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to run http server")
		panic(err)
	}
}
