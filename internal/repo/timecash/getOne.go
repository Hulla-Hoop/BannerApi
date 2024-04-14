package timecash

import "banner/internal/model"

func (c *timeCash) GetOne(reqId string, tag_id int, feature_id int, last bool, role bool) (model.Banner, error) {
	if !last {
		return c.db.GetOne(reqId, tag_id, feature_id, last, role)
	} else {
		key := [2]int{feature_id, tag_id}
		ban, ok := c.get(key)
		if ok {
			return ban, nil
		} else {
			ban, err := c.db.GetOne(reqId, tag_id, feature_id, last, role)
			if err != nil {
				return model.Banner{}, err
			}
			c.add(key, ban)
			return ban, nil
		}
	}
}
