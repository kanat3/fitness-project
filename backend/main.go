package main

import (
	"fitness-project/backend/api"
	"fitness-project/backend/internal/config"
	"fitness-project/backend/internal/storage"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config
	cfg := config.GetConfig()

	_ = cfg

	// connect to database

	var db storage.Storage

	err := db.ConnectToDB()
	//defer db.Close()

	if err != nil {
		fmt.Println(error.Error(err))
		return
	}

	// healcheck db
	db.HealCheck()

	// init router
	r := gin.Default()
	api.InitHandlers(r)

	// run server
	r.Run(":8080")
}
