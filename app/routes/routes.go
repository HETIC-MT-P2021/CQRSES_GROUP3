package routes

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/controllers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP3/app/middlewares"
	"github.com/gin-gonic/gin"
)

// Init initializes router with the following routes
func Init(r *gin.Engine) {
	r.POST("/register", controllers.CreateCustomer)
	r.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.GET("/search/:id", controllers.TestSearchService)
	api.GET("/articles/:id", controllers.GetArticleById)

	// Only for test without auth middleware
	api.POST("/articles", controllers.CreateArticle)
	api.POST("/articles/:id", controllers.UpdateArticle)

	api.Use(middlewares.CheckAuthorization)
	{
		api.GET("/", controllers.TestController)

		// V1
		v1 := api.Group("/v1")
		// Search
		v1.GET("/search", controllers.TestSearchService)

		// Articles
		v1.POST("/articles", controllers.CreateArticle)
		v1.POST("/articles/:id", controllers.UpdateArticle)
		v1.DELETE("/articles/:id", controllers.DeleteArticleById)

	}
}
