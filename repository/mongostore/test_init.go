package mongostore

import (
	"path/filepath"

	"github.com/cazicbor/BORIS_LEVEL_UP/config"
	"github.com/cazicbor/BORIS_LEVEL_UP/db"
)

var CONFIGPATH = filepath.Join("BORIS_LEVEL_UP/", "conf.json")

func initEnvTest() {
	config.SetUp("./../../conf.json")
	db.SetUpDB()
}
