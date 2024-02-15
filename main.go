package main

import (
	"sau-na/common"
	"sau-na/router"
)

func main() {
	common.LoadEnv()
	router.Router()
}
