package dataservice

import (
	"context"
	"errors"

	"auth/dataservice/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TokenRepo struct {
	DB *gorm.DB
}

func (repo *TokenRepo) SaveToken(token *models.Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := repo.DB.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"refresh_token"}),
		},
	).Create(&token)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TokenRepo) DeleteToken(token *models.Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := repo.DB.WithContext(ctx).Delete(&token)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TokenRepo) GetTokenByUserID(userID int) (*models.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var token models.Token

	result := repo.DB.WithContext(ctx).Where("user_id = ?", userID).First(&token)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &token, nil
}
