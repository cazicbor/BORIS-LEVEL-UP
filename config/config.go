package config

var config *Configuration

//Configuration configration struct
type Configuration struct {
	Server ServerConfiguration
	DB     DBConfiguration
}

type ServerConfiguration struct {
	Addr string `json:"adrr" mapstructure:"SERVER_ADDRESS"`
	Port string `json:"adrr" mapstructure:"SERVER_PORT"`
	Mode string `json:"adrr" mapstructure:"SERVER_ADDRESS"`
}

type DBConfiguration struct {
	Database     string `json:"database" mapstructure:"DB_NAME"`
	TestDatabase string `mapstructure:"DB_NAME_TEST"`
	MongoDBUser  string `json:"mongodbuser" mapstructure:"DB_USER"`
	MongoDBPwd   string `json:"mongodbpwd" mapstructure:"DB_USER_PASSWORD"`
	MongoDBHost  string `json:"mongodbhost" mapstructure:"DB_HOST"`
	MongoDBPort  string `json:"mongodbport" mapstructure:"DB_HOST_PORT"`
}

func GetConfig() *Configuration {
	return config
}

func (conf *Configuration) SetTestMod() {
	conf.DB.Database = conf.DB.TestDatabase
}
