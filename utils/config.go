package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Configuration struct {
	Logger         string
	MongoUrl       string
	MongoDbName    string
	HttpHostUrl    string
	LogFileName    string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	HandlerTimeout time.Duration
}

// Read configuration file and create Configuration
func Readconf(configFileName string) (Configuration, error) {
	configuration := Configuration{}
	if _, err := os.Stat(configFileName); err != nil {
		if os.IsNotExist(err) {
			return configuration, fmt.Errorf("Readconf(): config file does not exist!")
		}
	}
	file, _ := os.Open(configFileName)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		return configuration, fmt.Errorf("Readconf(): config file error!", err)
	}
	return configuration, nil
}

func (c Configuration) To_string() string {
	mes := ("Start application: http server=" + c.HttpHostUrl + "  mongo server=" + c.MongoUrl + "  logger mode=" + c.Logger)
	return mes
}
