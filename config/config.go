package config

import (
	"encoding/json"
	"log"
	"os"
)

var config *Configuration

type Configuration struct {
	Server ServerConfiguration `json:"server"`
	DB     DBConfiguration     `json:"db"`
}

type ServerConfiguration struct {
	Addr string `json:"adrr"`
	Port string `json:"port"`
	Mode string `json:"mode"` //= test ou pas
}

type DBConfiguration struct {
	Database     string `json:"database"`
	TestDatabase string `json:"testdatabase"`
	MongoDBUser  string `json:"mongodbuser"`
	MongoDBPwd   string `json:"mongodbpwd"`
	MongoDBHost  string `json:"mongodbhost"`
	MongoDBPort  string `json:"mongodbport"`
}

func SetUp(configFilePath string) {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Println(err)
	}
}

func GetConfig() *Configuration {
	return config
}
