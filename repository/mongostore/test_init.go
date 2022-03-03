package mongostore

import (
	"path/filepath"

	"github.com/cazicbor/BORIS_LEVEL_UP/config"
	"github.com/cazicbor/BORIS_LEVEL_UP/db"
)

var CONFIGPATH = filepath.Join("BORIS_LEVEL_UP/config", "config")

func initEnvTest() {
	config.SetUp(CONFIGPATH)
	conf := config.GetConfig()
	conf.SetTestMod()
	db.SetUpDB()
}
