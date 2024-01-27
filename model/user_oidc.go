package model

import (
	"log"

	"braces.dev/errtrace"
	"gorm.io/gorm"
)

type UserOIDC struct {
	Model
	UserID   UUID   `json:"user_id" gorm:"type:binary(16)"`
	User     User   `json:"user" gorm:"foreignKey:UserID"`
	Provider string `json:"provider" gorm:"size:20;not null"`
	Subject  string `json:"subject" gorm:"size:100;not null"`
}

func (UserOIDC) TableName() string {
	return "user_oidc"
}

func (u *UserOIDC) Create(db *gorm.DB) error {
	if err := db.Create(&u).Error; err != nil {
		log.Println("user_oidcの作成に失敗しました。")
		return errtrace.Wrap(err)
	}
	return nil
}

func (u *UserOIDC) FindByID(db *gorm.DB, id UUID) error {
	if err := db.Preload("User").Where("id = ?", id).First(&u).Error; err != nil {
		return errtrace.Wrap(err)
	}
	return nil
}

func (u *UserOIDC) FindByUserID(db *gorm.DB, userID UUID) error {
	if err := db.Where("user_id = ?", userID).First(&u).Error; err != nil {
		return errtrace.Wrap(err)
	}
	return nil
}

func (u *UserOIDC) FindBySubject(db *gorm.DB, subject string) error {
	if err := db.Where("subject = ?", subject).Preload("User").First(&u).Error; err != nil {
		return errtrace.Wrap(err)
	}
	return nil
}

func (u *UserOIDC) UserExistsForSubject(db *gorm.DB, subject string) bool {
	var exists bool
	err := db.
		Model(&u).
		Select("count(*) > 0").
		Where("subject = ?", subject).
		Find(&exists).Error
	if err != nil {
		return false
	}
	return exists
}
