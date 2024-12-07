package utils

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var Conf = newConf()

type conf struct {
	DB databaseConfig
}

type databaseConfig struct {
	DSN      string `toml:"dsn"`
}

func newConf() conf {
	var _conf conf
	confPath := "config/"
	env := os.Getenv("ENV")

	switch env {
	case "local":
		confPath += "conf.local.toml"
	case "dev":
		confPath += "conf.dev.toml"
	case "prod":
		confPath += "conf.prod.toml"
	default:
		confPath += "conf.local.toml"
		log.Println("ENV is invalid or ENV is not set. Defaulting to local configuration.")
	}
	log.Printf("Load configuration env=%s conf=%s", env, confPath)

	asset, err := os.Open(confPath)
	if err != nil {
		log.Fatalf("Failed to read configuration file. confPath: %s err: %s", confPath, err.Error())
	}
	defer asset.Close()

	decoder := toml.NewDecoder(asset)
	if _, err = decoder.Decode(&_conf); err != nil {
		log.Fatalf("[CONFIGURATION FILE LOAD ERROR] confPath: %s err: %s", confPath, err.Error())
	}

	return _conf
}