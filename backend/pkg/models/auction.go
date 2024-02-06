package models

import (
	"time"

	"github.com/google/uuid"
)

const auctionTable = "auction"

func (Auction) TableName() string {
	return auctionTable
}

type Auction struct {
	Id          *uuid.UUID `gorm:"column:id;->"`
	CreatorId   string     `gorm:"column:creatorid"`
	Title       string     `gorm:"column:title"`
	Description string     `gorm:"column:description"`
	StartBit    int64      `gorm:"column:startbit"`
	CurrentBit  int64      `gorm:"column:currentbit"`
	CreatedAt   time.Time  `gorm:"column:createdat;->"`
	UpdatedAt   time.Time  `gorm:"column:updatedat"`
}
