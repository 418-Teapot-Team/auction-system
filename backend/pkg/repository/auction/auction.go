package auction

import (
	"auction-system/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateAuction(auction *models.Auction) error
	GetAllAuctions() (auctions []models.Auction, err error)
	GetAuctionById(id string) (auction models.Auction, err error)
	GetAuctionsByUserId(id string) (auctions []models.Auction, err error)
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

func (db *repo) GetAllAuctions() (auctions []models.Auction, err error) {
	err = db.db.Preload("Images").Preload("User").Find(&auctions).Error
	if err != nil {
		return nil, err
	}
	return
}

func (db *repo) GetAuctionById(id string) (auction models.Auction, err error) {
	err = db.db.Model(&models.Auction{}).
		Where("id = ?", id).
		Preload("Images").
		Preload("User").Find(&auction).Error
	if err != nil {
		return
	}
	return
}

func (db *repo) GetAuctionsByUserId(id string) (auctions []models.Auction, err error) {
	err = db.db.Model(&models.Auction{}).
		Where("creatorid = ?", id).
		Preload("Images").Find(&auctions).Error
	if err != nil {
		return nil, err
	}
	return
}
