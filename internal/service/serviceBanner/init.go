package servicebanner

import (
	"banner/internal/config"
	"banner/internal/repo"

	"github.com/sirupsen/logrus"
)

type serviceBanner struct {
	logger *logrus.Logger
	db     repo.Repos
	cfg    *config.ConfigRemoteApi
}

func InitServiceBanner(logger *logrus.Logger, db repo.Repos) *serviceBanner {
	cfg := config.RemoteApi()
	return &serviceBanner{
		logger: logger,
		db:     db,
		cfg:    cfg,
	}
}

type ErrIncorrectData struct {
	msg string
}

func (e ErrIncorrectData) Error() string {
	return e.msg
}
