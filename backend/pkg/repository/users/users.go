package users

import (
	"auction-system/pkg/models"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserById(id string) (*models.User, error)
}

type repo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (db *repo) GetUserById(id string) (*models.User, error) {
	user := new(models.User)

	err := db.db.Where("id = ?", id).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
