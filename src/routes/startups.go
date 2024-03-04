package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/nilerajput91/go-gin-docker-aws/src/controllers"
)

func startupsGroupRouter(baseRouter *gin.RouterGroup) {
	startups := baseRouter.Group("/startups")

	startups.GET("/all", controllers.GetAllStartups)
	startups.GET("/get/:id", controllers.GetStartupByID)
	startups.POST("/create", controllers.CreateStartup)
	startups.PATCH("/update", controllers.UpdateStartup)
	startups.PUT("/update", controllers.UpdateStartup)
	startups.DELETE("/delete/:id", controllers.DeleteStartup)
}
