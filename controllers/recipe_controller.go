package controllers

import (
	"net/http"
	"recipe_backend/models"
	"recipe_backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := c.Get("userID")
	recipe.UserID = userID.(uint)
	if err := services.CreateRecipe(&recipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe created successfully"})
}

func GetRecipes(c *gin.Context) {
	recipes, err := services.GetRecipes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recipes)
}

func GetRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	recipe, err := services.GetRecipe(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}
	c.JSON(http.StatusOK, recipe)
}

func UpdateRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := c.Get("userID")
	recipe.UserID = userID.(uint)
	if err := services.UpdateRecipe(uint(id), &recipe); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe updated successfully"})
}

func DeleteRecipe(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("userID")
	if err := services.DeleteRecipe(uint(id), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}
