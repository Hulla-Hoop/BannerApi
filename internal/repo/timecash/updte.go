package timecash

import "banner/internal/model"

func (c *timeCash) Update(reqId string, b model.BannerDB, t model.Tags) error {

	return c.db.Update(reqId, b, t)
}
