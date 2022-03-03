package mongostore

import (
	"github.com/cazicbor/BORIS_LEVEL_UP/config"
	"github.com/cazicbor/BORIS_LEVEL_UP/db"
)

func initEnvTest() {
	conf := config.GetConfig()
	conf.SetTestMod()
	db.SetUpDB()
}
