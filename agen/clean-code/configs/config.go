package configs

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	Demo string `json:"demo" yaml:"demo" mapstructure:"demo"`
}

func LoadConfig(cPath ...string) *Config {
	v := viper.New()
	for _, path := range cPath {
		v.AddConfigPath(path)
	}
	v.AddConfigPath(".")
	v.SetConfigType("{{CONFIG_TYPE}}")
	v.SetConfigName("{{CONFIG_NAME}}")
	if err := v.ReadInConfig(); err != nil {
		log.Println("Error reading config file: ", err.Error())
	}
	err := v.Unmarshal(&Cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal config: ", err.Error())
		return nil
	}

	return Cfg

}
