package servicebanner

import (
	"banner/internal/model"
	"encoding/json"
)

func (c *serviceBanner) Update(reqId string, id string, ban []byte) error {

	var banner model.BannerHttp

	err := json.Unmarshal(ban, &banner)
	if err != nil {
		return err
	}

	c.logger.WithField("ServiceBanner.Update", reqId).Debug("Полученные данные -- ", banner, "----", id)

	idi, err := c.idCheckAndConvert(reqId, id)
	if err != nil {
		return err
	}

	banner.Banner_id = idi

	err = c.validate(reqId, banner)
	if err != nil {
		return err
	}

	b, t := banner.BFTOTagsAndBannerDB()

	err = c.db.Update(reqId, b, t)
	if err != nil {

		return err
	}

	return nil
}
