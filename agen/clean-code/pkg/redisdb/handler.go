package redisdb

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
)

type Config struct {
	UseSentinel  bool   `json:"sentinel" mapstructure:"sentinel"`
	MasterName   string `json:"master_name" mapstructure:"master_name"`
	URI          string `json:"uri"  mapstructure:"uri"`
	Username     string `json:"username" mapstructure:"username"`
	Password     string `json:"password" mapstructure:"password"`
	Database     int    `json:"database" mapstructure:"database"`
	MinIdleConns int    `json:"min_idle_connection" mapstructure:"min_idle_connection"`
	PoolSize     int    `json:"pool_size" mapstructure:"pool_size"`
}

type config struct {
	UseSentinel bool     `yaml:"sentinel"`
	MasterName  string   `yaml:"master_name"`
	Addrs       []string `yaml:"address"`

	Password string `yaml:"password"`
	Database int    `yaml:"database"`

	MaxRetries      int           `yaml:"max_retries"`
	MinRetryBackoff time.Duration `yaml:"min_retry_backoff"`
	MaxRetryBackoff time.Duration `yaml:"max_retry_backoff"`

	DialTimeout  time.Duration `yaml:"dial_timeout"`
	ReadTimeout  time.Duration `yaml:"read_timeout"`
	WriteTimeout time.Duration `yaml:"write_timeout"`

	PoolSize           int           `yaml:"pool_size"`
	MinIdleConns       int           `yaml:"min_idle_conns"`
	MaxConnAge         time.Duration `yaml:"max_conn_age"`
	PoolTimeout        time.Duration `yaml:"pool_timeout"`
	IdleTimeout        time.Duration `yaml:"idle_timeout"`
	IdleCheckFrequency time.Duration `yaml:"idle_check_frequency"`
}

// Parse handles config for both standalone server and sentinel
func (c *config) Parse() (*redis.Options, *redis.FailoverOptions, error) {
	if c.UseSentinel {
		if c.MasterName == "" {
			return nil, nil, errors.New("pkg/redis: invalid config, must provide master name when using sentinel")
		}
		if len(c.Addrs) == 0 {
			return nil, nil, errors.New("pkg/redis: invalid config, must provide address when using sentinel")
		}

	} else {
		if len(c.Addrs) != 1 {
			return nil, nil, errors.New("pkg/redis: invalid config, must provide exactly 1 address when using standalone")
		}
	}

	if !c.UseSentinel {
		opts, err := redis.ParseURL(c.Addrs[0])
		if err != nil {
			return nil, nil, err
		}
		if len(c.Password) > 0 && len(opts.Password) == 0 {
			opts.Password = c.Password
		}
		if opts.DB == 0 && c.Database > 0 {
			opts.DB = c.Database
		}

		opts.PoolSize = 1000

		return opts, nil, err
	}

	opts := &redis.FailoverOptions{
		MasterName:         c.MasterName,
		SentinelAddrs:      c.Addrs,
		OnConnect:          nil,
		Password:           c.Password,
		DB:                 c.Database,
		MaxRetries:         c.MaxRetries,
		MinRetryBackoff:    c.MinRetryBackoff,
		MaxRetryBackoff:    c.MaxRetryBackoff,
		DialTimeout:        c.DialTimeout,
		ReadTimeout:        c.ReadTimeout,
		WriteTimeout:       c.WriteTimeout,
		PoolSize:           c.PoolSize,
		MinIdleConns:       c.MinIdleConns,
		MaxConnAge:         c.MaxConnAge,
		PoolTimeout:        c.PoolTimeout,
		IdleTimeout:        c.IdleTimeout,
		IdleCheckFrequency: c.IdleCheckFrequency,
		TLSConfig:          nil,
	}
	return nil, opts, nil
}

func makeCfg(cfg *Config) *config {
	return &config{
		UseSentinel:  cfg.UseSentinel,
		MasterName:   cfg.MasterName,
		Addrs:        strings.Split(cfg.URI, " "),
		Password:     cfg.Password,
		Database:     cfg.Database,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		PoolSize:     1000,
		MinIdleConns: 50,
	}
}

func New(cfgEnv Config) (*redis.Client, error) {

	cfg := makeCfg(&cfgEnv)

	opts1, opts2, err := cfg.Parse()
	if err != nil {
		return nil, err
	}

	var client *redis.Client
	if opts1 != nil {
		client = redis.NewClient(opts1)
	} else {
		client = redis.NewFailoverClient(opts2)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err = client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil
}
