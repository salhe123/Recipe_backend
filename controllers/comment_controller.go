package controllers

import (
	"net/http"
	"recipe_backend/models"
	"recipe_backend/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentRecipe(c *gin.Context) {
	recipeID, _ := strconv.Atoi(c.Param("id"))
	userID, _ := c.Get("userID")
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.RecipeID = uint(recipeID)
	comment.UserID = userID.(uint)
	if err := services.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}
