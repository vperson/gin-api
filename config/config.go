package config

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	Name          string              `yaml:"name"`
	Log           LogConfig           `yaml:"log"`
	Server        ServerConfig        `yaml:"server"`
	DB            DBConfig            `yaml:"db"`
	MetricsServer MetricsServerConfig `yaml:"metricsServer"`
	Cache         CacheConfig         `yaml:"cache"`
}

func (c *Config) SetDefault() {
	c.Name = "gin-api"
	c.Log.SetDefault()
	c.Server.SetDefault()
	c.DB.SetDefault()
	c.MetricsServer.SetDefault()
	c.Cache.SetDefault()
}

var config *Config

func New(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	config = &Config{}
	config.SetDefault()

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func Get() *Config {
	return config
}
