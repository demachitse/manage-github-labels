package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"github.com/spf13/viper"
)

var Data Config

type Config struct {
	Email *string
	Token *string
    Labels []github.Label
}

func LoadConfig() error {
	cfd, err := os.UserConfigDir()
	if err != nil {
		return err
	}
	cfg := filepath.Join(cfd, "manage-github-labels")

	if _, err := os.Stat(cfg); os.IsNotExist(err) {
		return err
	}

	viper.AddConfigPath(cfg)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetDefault("email", "unknown")

	viper.Unmarshal(&Data)

	if Data.Token == nil {
		return errors.New("not found token")
	}

	return nil
}
