package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configs struct {
	Database
	Jwt
	Redis
}

var SystemConfigs *Configs

func InitConfig() {
	path, err := os.Getwd()
	// fmt.Println("page", path)
	if err != nil {
		log.Fatal(err)
	}
	viper.SetConfigName("configs")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/home/catcatproer/go/src/gin-vue-admin-framework-back")
	// viper.AddConfigPath(path)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading in config: %s", err.Error(), path)
	}

	var configs Configs

	err = viper.Unmarshal(&configs)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %s", err.Error())
	}
	// return &configs
	SystemConfigs = &configs
	// fmt.Println("test" + SystemConfigs.Database.Prefix)
}

// func init() {
// 	SystemConfigs = loadConfigs()
// }
