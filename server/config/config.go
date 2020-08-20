package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// A Config holds all configurable settings from a yml config file
type Config struct {
	ListenAddr       string
	ConnectionString string

	CORS struct {
		Debug  bool
		Origin string
	}
}

// New uses a file path stored in the CONFIG environment variable to populate the Config struct
func New() (*Config, error) {
	path := os.Getenv("CONFIG")

	if len(path) <= 0 {
		path = "config/config.yml"
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.Unmarshal(data, &cfg)

	if err != nil {
		return nil, err
	}

	//TODO: Validate Fields

	return cfg, nil
}
