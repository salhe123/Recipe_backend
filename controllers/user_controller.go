package controllers

import (
	"net/http"
	"recipe_backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserRecipes(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	recipes, err := services.GetRecipesByUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipes)
}
