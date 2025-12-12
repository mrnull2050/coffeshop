package config

import (


	"github.com/spf13/viper"
)


type Config struct{
	Service Service
	Postgres Postgres
}

type Service struct {
	Port    int
	RunMode string
}
type Postgres struct {
	Name     string
	Username string
	Password string
	Port     int
	Host     string

}

func LoadConfig()(*Config , error){
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err  := viper.ReadInConfig(); err != nil {
		return nil , err
	}

	cfg := &Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil ,err
	}
	return cfg , nil 
}
