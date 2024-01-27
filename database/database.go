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
	Env    string
}

func (DbConfig) LoadEnv() (*DbConfig, error) {

	env := os.Getenv("ENV")

	if env == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, errors.New("envの読み込みに失敗しました。")
		}
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
	conf, err := new(DbConfig).LoadEnv()
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%s:%s@(%s:%v)/?charset=utf8mb4&parseTime=true&loc=Local", conf.User, conf.Pass, conf.Host,
		conf.Port)

	if conf.Env == "dev" || conf.Env == "prod" {
		dsn += "?tls=true&interpolateParams=true"
	}

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
	createDatabaseForRoot()
	conf, err := new(DbConfig).LoadEnv()
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=true", conf.User, conf.Pass, conf.Host, conf.Port, conf.Dbname)

	if conf.Env == "dev" || conf.Env == "prod" {
		dsn += "?tls=true&interpolateParams=true"
	}

	db, err := gorm.Open(mysql.Open(dsn), GetGormConf())
	if err != nil {
		return nil, errtrace.Wrap(err)
	}
	return db, nil
}
