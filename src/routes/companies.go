package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/nilerajput91/go-gin-docker-aws/src/controllers"
)

func companiesGroupRouter(baseRouter *gin.RouterGroup) {

	companies := baseRouter.Group("/companies")
	companies.GET("/all", controllers.GetAllCompanies)
	companies.GET("/get/:id", controllers.GetCompanyByID)
	companies.POST("/create", controllers.CreateCompany)
	companies.PATCH("/update", controllers.UpdateCompany)
	companies.PUT("/update", controllers.UpdateCompany)
	companies.DELETE("/delete/:id", controllers.DeleteCompany)
}
