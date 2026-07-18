package repositories

import (
	"database/sql"
	"rent-car-project/config"
	"rent-car-project/models"
)

func CreateUser(user models.User) (int, error) {
	var userID int
	err := config.DB.QueryRow(
		"INSERT INTO users (name, email, password_hash, phone_number, role) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Name, user.Email, user.PasswordHash, user.PhoneNumber, user.Role,
	).Scan(&userID)
	return userID, err
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.QueryRow(
		"SELECT id, password_hash, role FROM users WHERE email = $1", 
		email,
	).Scan(&user.ID, &user.PasswordHash, &user.Role)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
