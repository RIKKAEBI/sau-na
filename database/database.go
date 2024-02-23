package database

import (
	"fmt"
	"sau-na/global"

	"braces.dev/errtrace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia%%2FTokyo",
		global.Env.DB_USER,
		global.Env.DB_PASS,
		global.Env.DB_HOST,
		global.Env.DB_NAME)

	db, err := gorm.Open(mysql.Open(dsn), gormConf())

	if err != nil {
		fmt.Print(dsn, "データベースに接続できませんでした。")
		errtrace.Wrap(err)
	}

	return db
}

func gormConf() *gorm.Config {
	return &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}
