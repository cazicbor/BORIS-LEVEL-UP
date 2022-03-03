package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *Configuration

type Configuration struct {
	Server ServerConfiguration
	DB     DBConfiguration
}

type ServerConfiguration struct {
	Addr string `json:"adrr" mapstructure:"SERVER_ADDRESS"`
	Port string `json:"adrr" mapstructure:"SERVER_PORT"`
	Mode string `json:"adrr" mapstructure:"SERVER_MODE"`
}

type DBConfiguration struct {
	Database     string `json:"database" mapstructure:"DB_NAME"`
	TestDatabase string `mapstructure:"DB_NAME_TEST"`
	MongoDBUser  string `json:"mongodbuser" mapstructure:"DB_USER"`
	MongoDBPwd   string `json:"mongodbpwd" mapstructure:"DB_USER_PASSWORD"`
	MongoDBHost  string `json:"mongodbhost" mapstructure:"DB_HOST"`
	MongoDBPort  string `json:"mongodbport" mapstructure:"DB_HOST_PORT"`
}

func SetUp(configFilePath string) {
	viper.AddConfigPath(configFilePath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("configuration not found")
		} else {
			log.Fatalf("%s", err)
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func GetConfig() *Configuration {
	return config
}

func (conf *Configuration) SetTestMod() {
	conf.DB.Database = conf.DB.TestDatabase
}
