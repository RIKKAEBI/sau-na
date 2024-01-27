package database

import (
	"sau-na/model"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	_ = db.AutoMigrate(&model.User{})
	_ = db.AutoMigrate(&model.UserOIDC{})
}
