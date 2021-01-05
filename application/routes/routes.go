package routes

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/controllers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/application/middlewares"
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

		// V1
		v1 := api.Group("/v1")
		v1.GET("/search", controllers.TestSearchService)
	}
	//test
	r.GET("/search", controllers.TestSearchService)
}
