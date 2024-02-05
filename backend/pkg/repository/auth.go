package repository

import "gorm.io/gorm"

type AuthRepository interface {
}

type repo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repo{db: db}
}
