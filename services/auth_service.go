package services

import (
	"errors"
	"rent-car-project/models"
	"rent-car-project/repositories"
	"rent-car-project/utils"
	"strings"
)

func RegisterUser(name, email, password, phone, roleInput string) (int, error) {
	role := "customer"
	if strings.ToLower(roleInput) == "owner" {
		role = "owner"
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return 0, errors.New("failed to hash password")
	}

	user := models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
		PhoneNumber:  phone,
		Role:         role,
	}

	userID, err := repositories.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return 0, errors.New("email already exists")
		}
		return 0, err
	}

	return userID, nil
}

func LoginUser(email, password string) (string, string, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return "", "", errors.New("database error")
	}
	if user == nil {
		return "", "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", "", errors.New("failed to generate token")
	}

	return token, user.Role, nil
}
