package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Name string   `mapstructure:"name"`
	List []string `mapstructure:"list"`
	Map  map[string]struct {
		Key string `mapstructure:"key"`
		URL string `mapstructure:"url"`
	} `mapstructure:"map"`
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	} else {
		fmt.Println("use config file", "path", viper.ConfigFileUsed())
	}

	viper.SetEnvPrefix("mf")
	viper.AutomaticEnv()

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
}
