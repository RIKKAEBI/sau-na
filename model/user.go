package model

import (
	"gorm.io/gorm"
)

type User struct {
	Model
	Name string `json:"name" gorm:"type:varchar(255);unique"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) Create(db *gorm.DB) (*User, error) {
	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
