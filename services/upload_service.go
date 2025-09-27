package services

import (
	"path/filepath"
	"recipe_backend/models"

	"github.com/gin-gonic/gin"
)

func UploadImages(c *gin.Context, recipeID uint) ([]models.Image, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	files := form.File["images"]
	featured := form.Value["featured"][0]

	var images []models.Image
	for _, file := range files {
		filename := filepath.Base(file.Filename)
		path := "uploads/" + filename
		if err := c.SaveUploadedFile(file, path); err != nil {
			return nil, err
		}
		isFeatured := filename == featured
		image := models.Image{
			RecipeID:   recipeID,
			URL:        path,
			IsFeatured: isFeatured,
		}
		if err := DB.Create(&image).Error; err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return images, nil
}
