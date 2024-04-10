package servicebanner

import (
	"banner/internal/model"
	"regexp"
)

func (c *serviceBanner) Filter(reqId string, tegID string, featureID string, limit string, offset string) ([]model.BannerHttp, error) {

	c.logger.WithField("ServiceBanner.Filter", reqId).Debug("Полученные данные teg -- ", tegID, "feature --", featureID)

	filter := c.filterValidate(reqId, tegID, featureID, limit, offset)

	return c.db.Filter(reqId, filter)
}

func (c *serviceBanner) filterValidate(reqId string, tegID string, featureID string, limit string, offset string) map[string]string {
	filter := make(map[string]string)

	if limit != "" {
		match, err := regexp.MatchString(`^[0-9]*$`, limit)
		if err != nil {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", limit)
		} else if !match {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", limit)
		} else {
			filter["limit"] = limit
		}
	}

	if offset != "" {
		match, err := regexp.MatchString(`^[0-9]*$`, offset)
		if err != nil {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", offset)
		} else if !match {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", offset)
		} else {
			filter["offset"] = offset
		}
	}

	if tegID != "" {
		match, err := regexp.MatchString(`^(?:1000|\d{1,3})$`, tegID)
		if err != nil {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", tegID)
		} else if !match {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", tegID)
		} else {
			filter["tag"] = tegID
		}
	}

	if featureID != "" {
		match, err := regexp.MatchString(`^(?:1000|\d{1,3})$`, featureID)
		if err != nil {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", featureID)
		} else if !match {
			c.logger.WithField("carCatalog.checkFilter", reqId).Error("некорректные данные ", featureID)
		} else {
			filter["feature"] = featureID
		}
	}
	return filter
}
