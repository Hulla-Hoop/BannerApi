package psql

import (
	"banner/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

type psql struct {
	dB     *sql.DB
	logger *logrus.Logger
}

func InitDb(logger *logrus.Logger) (*psql, error) {
	config := config.DbNew()

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=%s", config.Host, config.User, config.DBName, config.Password, config.Port, config.SSLMode)

	dB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	dB.SetMaxIdleConns(25)

	err = goose.Up(dB, "migration")
	if err != nil {
		return nil,
			fmt.Errorf("--- Ошибка миграции:%s", err)
	}
	return &psql{
		dB:     dB,
		logger: logger,
	}, nil

}

type ErrNotFound struct {
	msg string
}

func (e ErrNotFound) Error() string {
	return e.msg
}
