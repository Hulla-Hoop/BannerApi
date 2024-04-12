package midllware

import "github.com/sirupsen/logrus"

type midllware struct {
	L *logrus.Logger
}

func Init(l *logrus.Logger) *midllware {
	return &midllware{
		L: l,
	}
}
