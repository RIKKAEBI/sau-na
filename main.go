package main

import (
	"sau-na/common"
	"sau-na/conf"
	"sau-na/database"
	"sau-na/global"
	"sau-na/router"
)

func main() {
	global.Env = common.LoadEnv()
	global.DB = database.Connect()
	global.GoogleAuthConf = conf.GoogleAuthInit()

	router.Router()
}
