package configuration

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

const (
	CONFIG_FILENAME = "config.json"
)

var Config = Configuration{}

type Configuration struct {
	Cameras          []Camera
	DiscoveryAddress string
}

type Camera struct {
	Url           string
	Username      string
	Password      string
	AdminUsername string
	AdminPassword string
}

func Load() error {

	configFile, err := os.Open(CONFIG_FILENAME)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	decoder.Decode(&Config)

	//log.Println(Config)
	return nil
}

func Save() error {

	configFile, err := os.Create(CONFIG_FILENAME)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer configFile.Close()

	jsonConfig, err := json.Marshal(Config)

	if err != nil {
		log.Fatal(err)
		return err
	}

	n, err := io.WriteString(configFile, string(jsonConfig))

	if err != nil {
		log.Fatal(string(n), err.Error())
		return err
	}

	return nil
}
