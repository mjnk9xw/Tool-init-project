package configs

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Cfg *Config

type Config struct {
	Demo string `json:"demo" yaml:"demo" mapstructure:"demo"`
}

func LoadConfig(cPath ...string) *Config {
	v := viper.NewWithOptions(viper.KeyDelimiter("__"))
	customConfigPath := "."
	if len(cPath) > 0 {
		customConfigPath = cPath[0]
	}

	v.SetConfigType("json")
	defaultConfig, _ := json.Marshal(Cfg)
	err := v.ReadConfig(bytes.NewBuffer(defaultConfig))
	if err != nil {
		log.Fatal("Failed to read viper config: ", err.Error())
		return nil
	}

	v.SetConfigType("{{CONFIG_TYPE}}")
	v.SetConfigFile("{{CONFIG_NAME}}")
	if len(cPath) > 0 {
		v.SetConfigName(".env")
	}
	v.AddConfigPath(customConfigPath)
	if err := v.ReadInConfig(); err != nil {
		log.Println("Error reading config file: ", err.Error())
	}
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()
	err = v.Unmarshal(&Cfg)
	if err != nil {
		log.Fatal("Failed to unmarshal config: ", err.Error())
		return nil
	}

	return Cfg

}
