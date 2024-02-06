package models

const imagesTable = "auctioncontent"

func (Images) TableName() string {
	return imagesTable
}

type Images struct {
	Id          int64  `gorm:"column:id;->" json:"-"`
	AuctionId   string `gorm:"column:auctionid" json:"-"`
	DownloadUrl string `gorm:"column:downloadurl" json:"base64"`
}
