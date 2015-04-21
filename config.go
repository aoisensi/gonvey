package main

import "github.com/BurntSushi/toml"

type HTTPConfig struct {
	Public string
}

var Config struct {
	HTTP HTTPConfig
}

func LoadConfig() error {
	_, err := toml.DecodeFile(FlagConfig, &Config)
	return err
}
