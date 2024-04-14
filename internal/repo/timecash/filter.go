package timecash

import "banner/internal/model"

func (c *timeCash) Filter(reqId string, filter map[string]string) ([]model.BannerHttp, error) {
	return c.db.Filter(reqId, filter)
}
