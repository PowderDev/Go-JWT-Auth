package dataservice

import (
	"context"
	"errors"
	"time"

	"auth/dataservice/models"
	"gorm.io/gorm"
)

const dbTimeout = time.Second * 3

type UserRepo struct {
	DB *gorm.DB
}

func (repo *UserRepo) GetByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var user models.User

	result := repo.DB.WithContext(ctx).Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepo) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := repo.DB.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
