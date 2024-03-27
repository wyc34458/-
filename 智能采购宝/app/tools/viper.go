package tools

import (
	"github.com/spf13/viper"
)

type Config struct {
	MySql struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
	} `yaml:"mysql"`
	Redis struct {
		Address string `yaml:"address"`
	} `yaml:"redis"`
}

var Configs Config

func LoadConfig() {
	v := viper.New()
	v.SetConfigFile("D:/go/bysj1/app/tools/config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		return
	}
	err = v.Unmarshal(&Configs)
	if err != nil {
		return
	}
}
