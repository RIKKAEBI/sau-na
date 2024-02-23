package common

import (
	"fmt"

	caarlos0_env "github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Env struct {
	STAGE                string `env:"STAGE"`
	PORTOCOL             string `env:"PORTOCOL,required"`
	DOMAIN               string `env:"DOMAIN,required"`
	URL                  string `env:"URL"`
	GOOGLE_CLIENT_ID     string `env:"GOOGLE_CLIENT_ID,required"`
	GOOGLE_CLIENT_SECRET string `env:"GOOGLE_CLIENT_SECRET,required"`
	DB_HOST              string `env:"DB_HOST,required"`
	DB_USER              string `env:"DB_USER,required"`
	DB_PASS              string `env:"DB_PASS,required"`
	DB_NAME              string `env:"DB_NAME,required"`
}

// 環境変数を読み込む関数
func LoadEnv() Env {
	godotenv.Load(".env")

	var env Env
	if err := caarlos0_env.Parse(&env); err != nil {
		fmt.Println("必要な環境変数がありません", err)
	}

	env.URL = env.PORTOCOL + env.DOMAIN

	return env
}
