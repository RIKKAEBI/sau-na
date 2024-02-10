package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// 環境変数を読み込む関数
func LoadGoogleOAuthEnv() (GOOGLE_CLIENT_ID string, GOOGLE_CLIENT_SECRET string) {
	GOOGLE_CLIENT_ID, GOOGLE_CLIENT_ID_check := os.LookupEnv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET, GOOGLE_CLIENT_SECRET_check := os.LookupEnv("GOOGLE_CLIENT_SECRET")

	if !GOOGLE_CLIENT_ID_check || !GOOGLE_CLIENT_SECRET_check {
		godotenv.Load(".env")
		GOOGLE_CLIENT_ID, GOOGLE_CLIENT_ID_check = os.LookupEnv("GOOGLE_CLIENT_ID")
		GOOGLE_CLIENT_SECRET, GOOGLE_CLIENT_SECRET_check = os.LookupEnv("GOOGLE_CLIENT_SECRET")

		if !GOOGLE_CLIENT_ID_check || GOOGLE_CLIENT_SECRET_check {
			fmt.Printf("環境変数が読み込めてないからビルドを止めてあげたい")
		}
	}

	return GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET
}
