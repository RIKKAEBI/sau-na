package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// 環境変数を読み込む関数
func LoadEnv() {
	godotenv.Load(".env")
	arr := []string{"ORIGIN", "URL", "GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET"}

	for i := 0; i < len(arr); i++ {
		_, check := os.LookupEnv(arr[i])

		if !check {
			fmt.Printf("必要な環境変数がありません")
		}
	}
}

var ORIGIN = os.Getenv("ORIGIN")
var URL = os.Getenv("URL")
var GOOGLE_CLIENT_ID = os.Getenv("GOOGLE_CLIENT_ID")
var GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
