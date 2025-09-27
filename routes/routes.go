package routes

import (
	"recipe_backend/controllers"
	"recipe_backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	api := r.Group("/api").Use(middleware.AuthMiddleware())
	{
		api.POST("/recipes", controllers.CreateRecipe)
		api.GET("/recipes", controllers.GetRecipes)
		api.GET("/recipes/:id", controllers.GetRecipe)
		api.PUT("/recipes/:id", controllers.UpdateRecipe)
		api.DELETE("/recipes/:id", controllers.DeleteRecipe)
		api.GET("/categories", controllers.GetCategories)
		api.GET("/users/:id/recipes", controllers.GetUserRecipes)
		api.POST("/recipes/:id/like", controllers.LikeRecipe)
		api.POST("/recipes/:id/bookmark", controllers.BookmarkRecipe)
		api.POST("/recipes/:id/comment", controllers.CommentRecipe)
		api.POST("/recipes/:id/rating", controllers.RateRecipe)
		api.GET("/search", controllers.SearchRecipes)
	}
}
