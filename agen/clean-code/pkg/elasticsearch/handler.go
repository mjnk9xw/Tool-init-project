package elasticsearch

import (
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

type Config struct {
	ElasticSearchUrl      []string `json:"es_url" mapstructure:"es_url"`
	ElasticSearchUserName string   `json:"es_username" mapstructure:"es_username"`
	ElasticSearchPassword string   `json:"es_pass" mapstructure:"es_pass"`
}

func New(cfg *Config) (*elasticsearch.Client, error) {

	cfgElastic := elasticsearch.Config{
		Addresses:           cfg.ElasticSearchUrl,
		Username:            cfg.ElasticSearchUserName,
		Password:            cfg.ElasticSearchPassword,
		CompressRequestBody: true,
		EnableMetrics:       true,
		EnableDebugLogger:   false,
		DisableMetaHeader:   false,
	}
	es, err := elasticsearch.NewClient(cfgElastic)
	if err != nil {
		return es, fmt.Errorf("error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		return es, fmt.Errorf("error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
	return es, nil
}
