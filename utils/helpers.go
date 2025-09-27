package utils

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, err := regexp.MatchString(pattern, email)
	if err != nil || !matched {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasUpper || !hasLower || !hasNumber {
		return errors.New("password must contain uppercase, lowercase, and number")
	}
	return nil
}

func ValidateRecipeTitle(title string) error {
	if len(title) < 3 || len(title) > 100 {
		return errors.New("title must be between 3 and 100 characters")
	}
	return nil
}

func ValidatePrepTime(prepTime int) error {
	if prepTime <= 0 {
		return errors.New("preparation time must be greater than 0")
	}
	return nil
}

func ValidateComment(content string) error {
	if len(content) < 1 || len(content) > 500 {
		return errors.New("comment must be between 1 and 500 characters")
	}
	return nil
}

func ValidateRating(score int) error {
	if score < 1 || score > 5 {
		return errors.New("rating must be between 1 and 5")
	}
	return nil
}
