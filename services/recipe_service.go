package services

import (
	"errors"
	"recipe_backend/models"
	"strconv"
)

func CreateRecipe(recipe *models.Recipe) error {
	return DB.Create(recipe).Error
}

func GetRecipes() ([]models.Recipe, error) {
	var recipes []models.Recipe
	err := DB.Preload("Category").Preload("Images").Preload("Steps").Preload("Ingredients").Find(&recipes).Error
	return recipes, err
}

func GetRecipe(id uint) (models.Recipe, error) {
	var recipe models.Recipe
	err := DB.Preload("Category").Preload("Images").Preload("Steps").Preload("Ingredients").First(&recipe, id).Error
	return recipe, err
}

func UpdateRecipe(id uint, updatedRecipe *models.Recipe) error {
	var recipe models.Recipe
	if err := DB.First(&recipe, id).Error; err != nil {
		return err
	}
	recipe.Title = updatedRecipe.Title
	recipe.Description = updatedRecipe.Description
	recipe.PrepTime = updatedRecipe.PrepTime
	recipe.CategoryID = updatedRecipe.CategoryID
	return DB.Save(&recipe).Error
}

func DeleteRecipe(id, userID uint) error {
	var recipe models.Recipe
	if err := DB.First(&recipe, id).Error; err != nil {
		return err
	}
	if recipe.UserID != userID {
		return errors.New("unauthorized")
	}
	return DB.Delete(&recipe).Error
}

func GetRecipesByCategory(categoryID uint) ([]models.Recipe, error) {
	var recipes []models.Recipe
	err := DB.Where("category_id = ?", categoryID).Preload("Category").Preload("Images").Find(&recipes).Error
	return recipes, err
}

func GetRecipesByUser(userID uint) ([]models.Recipe, error) {
	var recipes []models.Recipe
	err := DB.Where("user_id = ?", userID).Preload("Category").Preload("Images").Find(&recipes).Error
	return recipes, err
}

func SearchRecipes(title, ingredient, prepTime string) ([]models.Recipe, error) {
	var recipes []models.Recipe
	query := DB.Preload("Category").Preload("Images").Preload("Steps").Preload("Ingredients")

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if ingredient != "" {
		query = query.Joins("JOIN ingredients ON ingredients.recipe_id = recipes.id").
			Where("ingredients.name LIKE ?", "%"+ingredient+"%")
	}
	if prepTime != "" {
		if pt, err := strconv.Atoi(prepTime); err == nil {
			query = query.Where("prep_time <= ?", pt)
		}
	}

	err := query.Find(&recipes).Error
	return recipes, err
}

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	err := DB.Find(&categories).Error
	return categories, err
}

func LikeRecipe(like *models.Like) error {
	var existing models.Like
	if err := DB.Where("user_id = ? AND recipe_id = ?", like.UserID, like.RecipeID).First(&existing).Error; err == nil {
		return errors.New("recipe already liked")
	}
	return DB.Create(like).Error
}

func BookmarkRecipe(bookmark *models.Bookmark) error {
	var existing models.Bookmark
	if err := DB.Where("user_id = ? AND recipe_id = ?", bookmark.UserID, bookmark.RecipeID).First(&existing).Error; err == nil {
		return errors.New("recipe already bookmarked")
	}
	return DB.Create(bookmark).Error
}

func CreateComment(comment *models.Comment) error {
	return DB.Create(comment).Error
}

func RateRecipe(rating *models.Rating) error {
	if rating.Score < 1 || rating.Score > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	var existing models.Rating
	if err := DB.Where("user_id = ? AND recipe_id = ?", rating.UserID, rating.RecipeID).First(&existing).Error; err == nil {
		return errors.New("recipe already rated")
	}
	return DB.Create(rating).Error
}
