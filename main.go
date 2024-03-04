package main

import (
	"github.com/nilerajput91/go-gin-docker-aws/src/middlewares"
	"github.com/nilerajput91/go-gin-docker-aws/src/models"
	"github.com/nilerajput91/go-gin-docker-aws/src/routes"
	"github.com/nilerajput91/go-gin-docker-aws/src/utils"
)

func main() {
	utils.LoadEnv()
	router := routes.SetupRoutes()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	middlewares.RegisterMiddlewares(router)
	router.Run(":8113")

}
