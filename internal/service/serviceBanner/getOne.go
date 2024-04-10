package servicebanner

import (
	"banner/internal/model"
	"strconv"
)

func (c *serviceBanner) GetOne(reqId string, tegID string, featureID string, las string) (model.Banner, error) {

	last, err := strconv.ParseBool(las)
	if err != nil {
		return model.Banner{}, err
	}

	c.logger.WithField("ServiceBanner.GetOne", reqId).Debug("Полученные данные teg -- ", tegID, "feature --", featureID, "last --", last)

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
