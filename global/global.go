package global

import (
	"sau-na/common"

	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

var Env common.Env

var DB *gorm.DB

var GoogleAuthConf *oauth2.Config
