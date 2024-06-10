package configs

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Configs struct {
	Database
}

var SystemConfigs *Configs

func loadConfigs() *Configs {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigName("configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath(path)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading in config: %s", err.Error(), path)
	}

	var configs Configs

	err = viper.Unmarshal(&configs)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %s", err.Error())
	}
	return &configs
}

func init() {
	SystemConfigs = loadConfigs()
}
