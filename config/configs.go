package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Path          string `json:"path"`
	ProjectName   string `json:"project_name"`
	Type          string `json:"type"`
	Model         string `json:"model"`
	Config        string `json:"config"`
	Db            string `json:"db"`
	Entities      string `json:"entities"`
	Usecase       string `json:"usecase"`
	DockerFile    bool   `json:"docker-file"`
	DockerCompose bool   `json:"docker-compose"`
}

func LoadConfig() *Config {
	var cfg Config
	raw, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("Error occured while reading config")
		return nil
	}
	json.Unmarshal(raw, &cfg)
	return &cfg
}
