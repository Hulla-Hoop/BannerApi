package timecash

import (
	"banner/internal/model"
	"banner/internal/repo"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	TTl       = 5 * time.Minute
	CleanTime = 5 * time.Minute
)

type value struct {
	Ban        model.Banner
	Expiration int64
}

type timeCash struct {
	db     repo.Repos
	logger *logrus.Logger
	Items  map[[2]int]value
	MU     sync.Mutex
}

func New(logger *logrus.Logger, db repo.Repos) *timeCash {
	c := &timeCash{
		logger: logger,
		Items:  make(map[[2]int]value),
		db:     db,
	}
	go c.remove()
	return c
}

func (c *timeCash) add(key [2]int, banner model.Banner) bool {
	c.MU.Lock()
	defer c.MU.Unlock()
	var val value
	val.Ban = banner
	val.Expiration = time.Now().Add(TTl).Unix()
	c.Items[key] = val
	c.logger.WithField("timeCash.add", key).Debug("значение добавлено в кеш data -- ", val)
	return true
}

func (c *timeCash) get(key [2]int) (model.Banner, bool) {
	c.MU.Lock()
	defer c.MU.Unlock()
	if data, ok := c.Items[key]; ok {
		c.logger.WithField("timeCash.get", key).Debug("значение взято из кеша data -- ", data)
		return data.Ban, true
	}
	return model.Banner{}, false
}

func (c *timeCash) remove() {

	for {
		time.Sleep(CleanTime)
		c.MU.Lock()
		now := time.Now().Unix()
		for k, v := range c.Items {
			if v.Expiration < now {
				c.logger.WithField("timeCash.remove", k).Debug("значение удалено из кеша data -- ", v)
				delete(c.Items, k)
			}
		}
		c.MU.Unlock()
	}

}
