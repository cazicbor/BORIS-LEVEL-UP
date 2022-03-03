package db

import (
	
	"github.com/cazicbor/BORIS_LEVEL_UP/config"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository/db"
	
func initEnvTest() {
	conf := config.GetConfig()
	conf.SetTestMod()
	db.SetUpDB()
}
