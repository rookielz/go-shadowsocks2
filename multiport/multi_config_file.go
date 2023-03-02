package multiport

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MultiConfig struct {
	Port     int
	Password string
}

type Config struct {
	Multi []MultiConfig
}

var config Config

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
}

func LoadFromYaml() {
	_ = viper.ReadInConfig()
	_ = viper.Unmarshal(&config)
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed: ", e.Name)
		for _, mulitConfig := range config.Multi {
			fmt.Printf("2：%d => %s \n", mulitConfig.Port, mulitConfig.Password)
		}
	})
	viper.WatchConfig()
	for _, mulitConfig := range config.Multi {
		fmt.Printf("1：%d => %s \n", mulitConfig.Port, mulitConfig.Password)
	}
}
