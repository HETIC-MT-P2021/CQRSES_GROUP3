package routes

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/middlewares"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/controllers"
	"github.com/gin-gonic/gin"
)

// Init initializes router with the following routes
func Init(r *gin.Engine) {
	r.POST("/register", controllers.CreateCustomer)
	r.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middlewares.CheckAuthorization)
	{
		api.GET("/", controllers.TestController)
	}
}
