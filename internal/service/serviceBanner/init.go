package servicebanner

import (
	"banner/internal/repo"

	"github.com/sirupsen/logrus"
)

type serviceBanner struct {
	logger *logrus.Logger
	db     repo.Repos
}

func InitServiceBanner(logger *logrus.Logger, db repo.Repos) *serviceBanner {
	return &serviceBanner{
		logger: logger,
		db:     db,
	}
}

type ErrIncorrectData struct {
	msg string
}

func (e ErrIncorrectData) Error() string {
	return e.msg
}
