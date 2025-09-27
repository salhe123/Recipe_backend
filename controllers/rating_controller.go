package controllers

import (
	"net/http"
	"recipe_backend/models"
	"recipe_backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RateRecipe(c *gin.Context) {
	recipeID, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("userID")
	var rating models.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rating.RecipeID = uint(recipeID)
	rating.UserID = userID.(uint)
	if err := services.RateRecipe(&rating); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe rated successfully"})
}
