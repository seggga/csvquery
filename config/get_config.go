package config

import (
	toml "github.com/pelletier/go-toml"
)

type ConfigType struct {
	Timeout int64
}

func GetConfig(filePath string) (*ConfigType, error) {

	// check file existance

	config, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	timeout := config.Get("csvquery.timeout").(int64)

	return &ConfigType{Timeout: timeout}, nil
}
