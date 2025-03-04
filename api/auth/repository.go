package auth

import (
	"personal-dashboard-backend/api/user/model"

	"gorm.io/gorm"
)

type AuthRepository struct { 
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (repo *AuthRepository) CreateUser(user *model.User) error {
	return repo.db.Create(user).Error
}

func (repo *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	return &user, err
}