package services

import (
	"errors"
	"rent-car-project/models"
	"rent-car-project/repositories"
)

func CreateCategory(name, description string) (int, error) {
	return repositories.CreateCategory(name, description)
}

func GetCategories() ([]models.Category, error) {
	return repositories.GetCategories()
}

func GetCategoryByID(id int) (*models.Category, error) {
	cat, err := repositories.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, errors.New("category not found")
	}
	return cat, nil
}

func UpdateCategory(id int, name, description string) error {
	cat, err := repositories.GetCategoryByID(id)
	if err != nil {
		return err
	}
	if cat == nil {
		return errors.New("category not found")
	}
	return repositories.UpdateCategory(id, name, description)
}

func DeleteCategory(id int) error {
	cat, err := repositories.GetCategoryByID(id)
	if err != nil {
		return err
	}
	if cat == nil {
		return errors.New("category not found")
	}
	return repositories.DeleteCategory(id)
}
