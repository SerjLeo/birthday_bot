package config

import (
	"github.com/spf13/viper"
)

const (
	defaultHTTPPort = "8000"
)

type (
	Config struct {
		SQLitePath string `yaml:"sqlitepath"`
	}

	BotConfig struct {
		Token string
	}
)

func InitConfig(configDir string) (*Config, error) {
	setDefaults()

	if err := parseFile(configDir); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseFile(directory) error {
	viper.AddConfigPath(directory)
	viper.SetConfigName("common")

	return viper.ReadInConfig()

}

func setDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
}
