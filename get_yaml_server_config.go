package get_yaml_server_config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type ServerConfig struct {
	Host    string  `yaml:"host" json:"host"`
	Port    int     `yaml:"port" json:"port"`
	Timeout float32 `yaml:"timeout" json:"timeout"`
}

func getConfig() (*ServerConfig, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}

	config := new(ServerConfig)
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
