package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/yuuki0310/reservation_api/config"
)

var Conf = newConf()

type conf struct {
	DB databaseConfig
	Cognito cognitoConfig
}

type databaseConfig struct {
	DSN string `toml:"dsn"`
	
}

type cognitoConfig struct {
	JWKSURL string `toml:"jwks_url"`
}

func newConf() conf {
	var _conf conf
	var confPath string
	env := os.Getenv("ENV")

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting working directory: %v\n", err)
	}
	fmt.Printf("Current working directory: %s\n", cwd)

	switch env {
	case "local":
		confPath = "local.toml"
	case "dev":
		confPath = "dev.toml"
	case "prod":
		confPath = "prod.toml"
	default:
		confPath += "local.toml"
		log.Println("ENV is invalid or ENV is not set. Defaulting to local configuration.")
	}
	log.Printf("Load configuration env=%s conf=%s", env, confPath)

	asset, err := config.Embed.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Failed to read configuration file. confPath: %s err: %s", confPath, err.Error())
	}

	_, err = toml.Decode(string(asset), &_conf)
	if err != nil {
		log.Fatalf("[CONFIGURATION FILE LOAD ERROR] confPath: %s err:%s", confPath, err.Error())
	}

	return _conf
}
