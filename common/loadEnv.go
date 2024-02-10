package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// 環境変数を読み込む関数
func LoadEnv() (origin string, URL string) {
	origin, originCheck := os.LookupEnv("ORIGIN")
	URL, URLCheck := os.LookupEnv("URL")

	if !originCheck || !URLCheck {
		godotenv.Load(".env")
		origin, originCheck = os.LookupEnv("ORIGIN")
		URL, URLCheck = os.LookupEnv("URL")

		if !originCheck || !URLCheck {
			fmt.Printf("環境変数が読み込めてないからビルドを止めてあげたい")
		}
	}

	return origin, URL
}
