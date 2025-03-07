package repository

import (
	"errors"
	"file-sharing/pkg/repository/interfaces"
	"file-sharing/pkg/utils/models"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

// constructor funciton

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		DB: DB,
	}
}

func (ur *userRepository) SignUp(user models.UserDetails) (models.UserDetailsResponse, error) {
	var userResponse models.UserDetailsResponse
	err := ur.DB.Raw("INSERT INTO users(name,email,username,phone,password)VALUES(?,?,?,?,?)RETURNING id,name,email,phone", user.Name, user.Email, user.Username, user.Phone, user.Password).Scan(&userResponse).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	return userResponse, nil
}
func (ur *userRepository) FindUserByEmail(user models.UserLogin) (models.UserSignInResponse, error) {
	var userResponse models.UserSignInResponse
	err := ur.DB.Raw("SELECT * FROM users WHERE email=? ", user.Email).Scan(&userResponse).Error
	if err != nil {
		return models.UserSignInResponse{}, errors.New("no user found")
	}
	return userResponse, nil
}
func (ur *userRepository) CheckUserAvailability(email string) bool {
	var userCount int

	err := ur.DB.Raw("SELECT COUNT(*) FROM users WHERE email=?", email).Scan(&userCount).Error
	if err != nil {
		return false
	}
	// if count greater than 0, user already exist
	return userCount > 0
}