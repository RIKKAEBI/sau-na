package main

import (
	"sau-na/common"
	"sau-na/router"
)

func main() {
	common.Env = common.LoadEnv()
	router.Router()
}
