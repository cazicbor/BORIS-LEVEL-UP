package mongostore

import (
	"github.com/cazicbor/BORIS_LEVEL_UP/config"
	"github.com/cazicbor/BORIS_LEVEL_UP/repository/mongostore"
)

func initEnvTest() {
	conf := config.GetConfig()
	conf.SetTestMod()
	mongostore.SetUpDB()
}
