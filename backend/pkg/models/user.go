package models

import "github.com/google/uuid"

const usersTable = "users"

func (User) TableName() string {
	return usersTable
}

type User struct {
	Id           *uuid.UUID `gorm:"column:id;->"`
	FullName     string     `gorm:"column:fullname;unique"`
	Username     string     `gorm:"column:username"`
	PasswordHash string     `gorm:"column:password"`
}
