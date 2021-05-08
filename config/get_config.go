package config

import (
	"errors"

	toml "github.com/pelletier/go-toml"
)

type ConfigType struct {
	Timeout  int64
	Graceful int64
}

var (
	errTime  error = errors.New("'timeout' has not been set in the config-file")
	errGrace error = errors.New("'graceful' has not been set in the config-file")
)

func GetConfig(filePath string) (*ConfigType, error) {

	// check file existance
	config, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	// obtain data
	timeout := config.Get("csvquery.timeout").(int64)
	grace := config.Get("csvquery.timeout").(int64)

	// check values
	if timeout == 0 {
		return nil, errTime
	}
	if grace == 0 {
		return nil, errGrace
	}

	return &ConfigType{Timeout: timeout, Graceful: grace}, nil
}
