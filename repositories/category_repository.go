package repositories

import (
	"database/sql"
	"rent-car-project/config"
	"rent-car-project/models"
)

func CreateCategory(name, description string) (int, error) {
	var id int
	err := config.DB.QueryRow(`
		INSERT INTO categories (name, description)
		VALUES ($1, $2) RETURNING id
	`, name, description).Scan(&id)
	return id, err
}

func GetCategories() ([]models.Category, error) {
	rows, err := config.DB.Query(`SELECT id, name, description, created_at, updated_at FROM categories ORDER BY id ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name, &cat.Description, &cat.CreatedAt, &cat.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}

func GetCategoryByID(id int) (*models.Category, error) {
	var cat models.Category
	err := config.DB.QueryRow(`SELECT id, name, description, created_at, updated_at FROM categories WHERE id = $1`, id).
		Scan(&cat.ID, &cat.Name, &cat.Description, &cat.CreatedAt, &cat.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &cat, nil
}

func UpdateCategory(id int, name, description string) error {
	_, err := config.DB.Exec(`
		UPDATE categories SET name = $1, description = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3
	`, name, description, id)
	return err
}

func DeleteCategory(id int) error {
	_, err := config.DB.Exec(`DELETE FROM categories WHERE id = $1`, id)
	return err
}
