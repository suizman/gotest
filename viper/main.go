package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigYAML struct {
	Config []ConfigMapping `mapstructure:"Mapping"`
}

type ConfigMapping struct {
	Country  string   `mapstructure:"Country"`
	Brands   []string `mapstructure:"Brands"`
	Subbrand []string `mapstructure:"Subbrand"`
}

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Error trying to read config file: %v", err)
		return
	}

	confMap := ConfigYAML{}
	viper.Unmarshal(&confMap)

	logrus.Infof("All settings from config: %v", viper.AllSettings())
	logrus.Infof("Get settings from mapping: %v", confMap)

}
