package timecash

import "banner/internal/model"

func (c *timeCash) Insert(reqId string, b model.BannerDB, t model.Tags) (int, error) {
	return c.db.Insert(reqId, b, t)
}
