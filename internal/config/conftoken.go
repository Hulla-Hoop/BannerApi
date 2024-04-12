package config

import (
	"os"
)

type configToken struct {
	SecretKey string
}

func TokenCFG() *configToken {

	return &configToken{
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}
