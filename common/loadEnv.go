package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// 環境変数を読み込む関数
func LoadEnv() string {
	origin, check := os.LookupEnv("ORIGIN")

	if !check {
		godotenv.Load(".env")
		origin, check = os.LookupEnv("ORIGIN")

		if !check {
			fmt.Printf("環境変数が読み込めてないからビルドを止めてあげたい")
		}
	}

	return origin
}
