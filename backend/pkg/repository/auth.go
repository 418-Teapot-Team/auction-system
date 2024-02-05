package repository

import (
	"errors"

	"auction-system/pkg/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	CreateUser(user *models.User) (err error)
	GetUserByUsername(username string) (*models.User, error)
}

type repo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repo{db: db}
}

func (db *repo) CreateUser(user *models.User) (err error) {
	tx := db.db.Begin()

	if err = tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (db *repo) GetUserByUsername(username string) (*models.User, error) {
	user := new(models.User)

	err := db.db.Where("username = ?", username).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
