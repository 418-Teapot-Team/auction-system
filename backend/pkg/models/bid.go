package models

import (
	"time"

	"github.com/google/uuid"
)

const bidsTable = "bids"

func (Bid) TableName() string {
	return bidsTable
}

type Bid struct {
	User User `gorm:"foreignKey:BidderId;" json:"bidder,omitempty"`
	Auction Auction `gorm:"foreignKey:AuctionId;" json:"auction,omitempty"`

	Id          *uuid.UUID `gorm:"column:id;->" json:"id"`
	BidderId   string     `gorm:"column:bidderid" json:"-"`
	AuctionId   string     `gorm:"column:auctionid" json:"-"`
	NewValue    int64      `gorm:"column:newvalue" json:"newValue"`
	CreatedAt   time.Time  `gorm:"column:createdat;->" json:"createdAt"`
}