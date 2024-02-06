package models

import "github.com/google/uuid"

const usersTable = "users"

func (User) TableName() string {
	return usersTable
}

type User struct {
	Id           *uuid.UUID `gorm:"column:id;->" json:"id"`
	FullName     string     `gorm:"column:fullname;unique" json:"fullName"`
	Username     string     `gorm:"column:username" json:"username"`
	PasswordHash string     `gorm:"column:password" json:"-"`
}
