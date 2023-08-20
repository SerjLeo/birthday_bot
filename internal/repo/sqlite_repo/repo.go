package sqlite_repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SerjLeo/birthday_bot/internal/models"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"time"
)

const (
	userTable = "user"
)

type SQLiteRepo struct {
	db *sql.DB
}

type Config struct {
	Path string
}

func New(config Config) (*SQLiteRepo, error) {
	db, err := sql.Open("sqlite3", config.Path)
	if err != nil {
		return nil, errors.Wrap(err, "connecting to sqlite")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "trying to ping database")
	}
	repo := &SQLiteRepo{db: db}
	err = repo.init()
	if err != nil {
		return nil, errors.Wrap(err, "initializing database")
	}
	return repo, nil
}

// init create tables if needed.
func (r *SQLiteRepo) init() error {
	q := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY, name TEXT UNIQUE, birthday TEXT)`, userTable)
	_, err := r.db.Exec(q)
	return err
}

// AddUser add user with birthday to database.
func (r *SQLiteRepo) AddUser(ctx context.Context, u *models.User) error {
	q := fmt.Sprintf(`INSERT INTO %s (name, birthday) VALUES (?, ?)`, userTable)
	if _, err := r.db.ExecContext(ctx, q, u.Name, u.Birthday); err != nil {
		return errors.Wrap(err, "saving user to db")
	}
	return nil
}

// GetUsersByMonth return list of users with birthday in given month.
func (r *SQLiteRepo) GetUsersByMonth(ctx context.Context, month time.Month) ([]*models.User, error) {
	return []*models.User{}, nil
}

// SearchUser return list of users with name due to search query.
func (r *SQLiteRepo) SearchUser(ctx context.Context, name string) ([]*models.User, error) {
	return []*models.User{}, nil
}

// DeleteUser delete certain user birthday record from database.
func (r *SQLiteRepo) DeleteUser(ctx context.Context, userId int64) error {
	return nil
}
