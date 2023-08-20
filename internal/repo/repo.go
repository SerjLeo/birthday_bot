package repo

import "github.com/SerjLeo/birthday_bot/internal/repo/sqlite_repo"

type Config struct {
	Path string
}

func New(config Config) (*sqlite_repo.SQLiteRepo, error) {
	return sqlite_repo.New(sqlite_repo.Config{
		Path: config.Path,
	})
}
