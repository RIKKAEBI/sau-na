package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"braces.dev/errtrace"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host   string
	User   string
	Pass   string
	Port   string
	Dbname string
}

func (DbConfig) LoadFromEnv() (*DbConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("envの読み込みに失敗しました。")
		return nil, errors.New("envの読み込みに失敗しました。")
	}
	conf := new(DbConfig)

	conf.Host = os.Getenv("DB_HOST")
	conf.User = os.Getenv("DB_USER")
	conf.Pass = os.Getenv("DB_PASS")
	conf.Port = os.Getenv("DB_PORT")
	conf.Dbname = os.Getenv("DB_NAME")
	return conf, nil
}

func GetGormConf() *gorm.Config {
	return &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

func createDatabaseForRoot() error {
	conf, err := new(DbConfig).LoadFromEnv()
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%v)/?charset=utf8mb4&parseTime=true&loc=Local&tls=true&interpolateParams=true", conf.User, conf.Pass, conf.Host, conf.Port)
	db, err := gorm.Open(mysql.Open(dsn), GetGormConf())
	if err != nil {
		return errtrace.Wrap(err)
	}

	db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", conf.Dbname))
	d, _ := db.DB()
	d.Close()
	return nil
}

func ConnectDatabase() (*gorm.DB, error) {
	conf, err := new(DbConfig).LoadFromEnv()
	dsn := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=true&tls=true&interpolateParams=true", conf.User, conf.Pass, conf.Host, conf.Port, conf.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), GetGormConf())
	if err != nil {
		return nil, errtrace.Wrap(err)
	}
	return db, nil
}
