package endpointsbaner

import (
	"banner/internal/service"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type endpoint struct {
	s      service.ServiceBanner
	logger *logrus.Logger
}

func Init(logger *logrus.Logger, s service.ServiceBanner) *endpoint {
	return &endpoint{
		s:      s,
		logger: logger,
	}
}

func (*endpoint) jsonError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
