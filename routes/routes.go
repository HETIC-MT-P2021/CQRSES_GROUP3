package routes

import (
	"github.com/edwinvautier/go-boilerplate/middlewares"
	"github.com/edwinvautier/go-boilerplate/controllers"
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
