package controllers

import (
	"net/http"
	"recipe_backend/models"
	"recipe_backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BookmarkRecipe(c *gin.Context) {
	recipeID, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("userID")
	bookmark := models.Bookmark{RecipeID: uint(recipeID), UserID: userID.(uint)}
	if err := services.BookmarkRecipe(&bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe bookmarked successfully"})
}
