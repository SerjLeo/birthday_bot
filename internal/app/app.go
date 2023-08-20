package app

import (
	"github.com/SerjLeo/birthday_bot/internal/config"
	"github.com/SerjLeo/birthday_bot/internal/repo"
)

func InitApp() error {
	cfg, err := config.InitConfig("config")
	if err != nil {
		return err
	}
	_, err = repo.New(repo.Config{Path: cfg.SQLitePath})
	if err != nil {
		return err
	}
}
