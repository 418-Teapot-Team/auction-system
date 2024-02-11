package bids

import (
	"auction-system/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateBid(bid *models.Bid) error
	GetAllBidsForAuction(auctionId string) (bids []models.Bid, err error)
}

type repo struct {
	db *gorm.DB
}

func NewBidsRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (db *repo) CreateBid(bid *models.Bid) error {
	tx := db.db.Begin()
	if err := tx.Clauses(clause.Returning{}).Select("*").Create(bid).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (db *repo) GetAllBidsForAuction(auctionId string) (bids []models.Bid, err error) {
	err = db.db.Model(&models.Bid{}).
		Where("auctionid = ?", auctionId).
		Preload("User").
		Find(&bids).Error
	if err != nil {
		return nil, err
	}
	return
}
