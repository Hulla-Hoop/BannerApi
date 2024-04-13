package servicebanner

import (
	"banner/internal/model"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

func (c *serviceBanner) GetOne(reqId string, tegID string, featureID string, las string, role bool) (model.Banner, error) {

	last, err := strconv.ParseBool(las)
	if err != nil {
		return model.Banner{}, err
	}

	c.logger.WithField("ServiceBanner.GetOne", reqId).Debug("Полученные данные teg -- ", tegID, "feature --", featureID, "last --", last)

	tegIDI, err := c.tegCheckAndConvert(reqId, tegID)
	if err != nil {
		return model.Banner{}, err
	}

	featureIDI, err := c.tegCheckAndConvert(reqId, featureID)
	if err != nil {
		return model.Banner{}, err
	}

	return c.db.GetOne(reqId, tegIDI, featureIDI, last, role)
}

func (c *serviceBanner) tegCheckAndConvert(reqId string, id string) (int, error) {
	pattern := `^(?:1000|\d{1,3})$`
	match, _ := regexp.MatchString(pattern, id)
	if !match {
		c.logger.WithField("ServiceBanner.idCheckAndConvert", reqId).Error("некорректные данные ", id)
		return -1, ErrIncorrectData{msg: fmt.Sprintf("данные %s некорректны", id)}
	}
	idi, err := strconv.Atoi(id)
	if err != nil {
		c.logger.WithField("ServiceBanner.idCheckAndConvert", reqId).Error("Ннеудалось преобразовать ", id)
		return -1, errors.New("неудалось преобразовать ")
	}

	return idi, nil
}
