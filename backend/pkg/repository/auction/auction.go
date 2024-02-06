package auction

import (
	"auction-system/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateAuction(auction *models.Auction) error
}

type repo struct {
	db *gorm.DB
}

func NewAuctionRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (db *repo) CreateAuction(auction *models.Auction) error {
	tx := db.db.Begin()
	if err := tx.Clauses(clause.Returning{}).Select("*").Create(auction).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
