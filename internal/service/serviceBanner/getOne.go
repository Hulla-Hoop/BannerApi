package servicebanner

import "banner/internal/model"

func (c *serviceBanner) GetOne(reqId string, tegID string, featureID string, last bool) (model.Banner, error) {

	c.logger.WithField("ServiceBanner.GetOne", reqId).Debug("Полученные данные teg -- ", tegID, "feature --", featureID)

	tegIDI, err := c.idCheckAndConvert(reqId, tegID)
	if err != nil {
		return model.Banner{}, err
	}

	featureIDI, err := c.idCheckAndConvert(reqId, featureID)
	if err != nil {
		return model.Banner{}, err
	}

	return c.db.GetOne(reqId, tegIDI, featureIDI, last)
}
