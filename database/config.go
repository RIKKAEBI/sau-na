package database

// "errors"
// "fmt"
// "os"

// "braces.dev/errtrace"
// "github.com/aws/aws-sdk-go-v2/config"
// "github.com/joho/godotenv"
// "gorm.io/driver/mysql"
// "gorm.io/gorm"

// type DbConfig struct {
// 	Host   string
// 	User   string
// 	Pass   string
// 	Port   string
// 	Dbname string
// }

// func (DbConfig) LoadFromEnv() error {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		return errors.New("envの読み込みに失敗しました。")
// 	}
// 	conf := new(DbConfig)

// 	conf.Host = os.Getenv("DB_HOST")
// 	conf.User = os.Getenv("DB_USER")
// 	conf.Pass = os.Getenv("DB_PASS")
// 	conf.Port = os.Getenv("DB_PORT")
// 	conf.Dbname = os.Getenv("DB_NAME")
// }

// func createDatabaseForRoot(config *config.Config) error {
// 	// conf := new(DbConfig).LoadFromEnv()
// 	dsn := fmt.Sprintf("%s:%s@(%s:%v)/?charset=utf8mb4&parseTime=true&loc=Local", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port)
// 	db, err := gorm.Open(mysql.Open(dsn), config)
// 	if err != nil {
// 		return nil, errtrace.Wrap(err)
// 	}

// 	db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", config.Database.Dbname))
// 	d, _ := db.DB()
// 	d.Close()
// 	return nil
// }

// func ConnectDatabaseForRoot(config *gorm.Config) (*gorm.DB, error) {
// 	createDatabaseForRoot(config)

// 	dsn := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8mb4&parseTime=true", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Dbname)
// 	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
// 	if err != nil {
// 		return nil, errtrace.Wrap(err)
// 	}
// }
