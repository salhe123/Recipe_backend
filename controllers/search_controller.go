package controllers

import (
	"net/http"
	"recipe_backend/services"

	"github.com/gin-gonic/gin"
)

func SearchRecipes(c *gin.Context) {
	title := c.Query("title")
	ingredient := c.Query("ingredient")
	prepTime := c.Query("prepTime")
	recipes, err := services.SearchRecipes(title, ingredient, prepTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipes)

}
