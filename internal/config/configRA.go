package config

import (
	"os"
)

type ConfigRemoteApi struct {
	Link string
}

func RemoteApi() *ConfigRemoteApi {

	return &ConfigRemoteApi{
		Link: os.Getenv("LINK"),
	}
}
