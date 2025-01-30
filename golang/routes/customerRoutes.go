package routes

import (
	"tech-test/controllers"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(r *gin.Engine) {
	customerGroup := r.Group("/customers")
	{
		customerGroup.GET("/", controllers.GetCustomers)
		customerGroup.GET("/:id", controllers.GetCustomerByID)
		customerGroup.POST("/", controllers.CreateCustomer)
		customerGroup.PUT("/:id", controllers.UpdateCustomer)
		customerGroup.DELETE("/:id", controllers.DeleteCustomer)
	}
}
