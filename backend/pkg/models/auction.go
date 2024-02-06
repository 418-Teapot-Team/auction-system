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
	User User `gorm:"foreignKey:CreatorId;" json:"creator"`

	Id          *uuid.UUID `gorm:"column:id;->" json:"id"`
	CreatorId   string     `gorm:"column:creatorid" json:"-"`
	Title       string     `gorm:"column:title" json:"title"`
	Description string     `gorm:"column:description" json:"description"`
	StartBit    int64      `gorm:"column:startbit" json:"startBit"`
	CurrentBit  int64      `gorm:"column:currentbit" json:"currentBit"`
	CreatedAt   time.Time  `gorm:"column:createdat;->" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"column:updatedat" json:"updatedAt"`
	Images      []Images   `gorm:"foreignKey:AuctionId;" json:"images"`
}
