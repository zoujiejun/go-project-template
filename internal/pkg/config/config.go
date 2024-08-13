package config

import (
	"github.com/spf13/viper"
)

func New(cfg string) (*Config, error) {
	var v = viper.New()
	v.SetConfigFile(cfg)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	var c = &Config{}
	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}

	return c, nil
}
