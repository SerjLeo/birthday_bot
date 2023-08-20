package service

import "github.com/pkg/errors"

type Repo interface {
	AddUser() error
}

type Service struct {
	repo *Repo
}

func New(repo *Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddUser() error {
	err := s.AddUser()
	if err != nil {
		return errors.Wrap(err, "add user to db")
	}
	return nil
}
