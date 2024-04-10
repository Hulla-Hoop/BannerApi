package servicebanner

import (
	"banner/internal/model"
	"encoding/json"
	"fmt"
	"regexp"
)

func (c *serviceBanner) Insert(reqId string, body []byte) (int, error) {

	var banner model.BannerHttp

	err := json.Unmarshal(body, &banner)
	if err != nil {
		return 0, err
	}

	c.logger.WithField("ServiceBanner.Insert", reqId).Debug("Полученные данные -- ", banner)

	err = c.validate(reqId, banner)
	if err != nil {
		return 0, err
	}

	b, t := banner.BFTOTagsAndBannerDB()

	return c.db.Insert(reqId, b, t)

}

func (c *serviceBanner) validate(reqId string, banner model.BannerHttp) error {

	if banner.Feature_id <= 0 {
		c.logger.WithField("ServiceBanner.validate", reqId).Error("некорректные данные ", banner)
		return ErrIncorrectData{msg: fmt.Sprintf("данные %d некорректны", banner.Feature_id)}
	}

	var errCount int
	for _, v := range banner.Tags_id {

		if v <= 0 {
			errCount++
			c.logger.WithField("ServiceBanner.validate", reqId).Error("некорректные тег -- ", v)
		}
	}
	if errCount == len(banner.Tags_id) {
		return ErrIncorrectData{msg: "все теги некорректны должен быть хотя бы один корректный тег"}
	}

	ok, err := regexp.MatchString(`^(https?|ftp):\/\/[^\s\/$.?#].[^\s]*$`, banner.Content.Url)
	if !ok || err != nil {
		c.logger.WithField("ServiceBanner.validate", reqId).Error("некорректные данные ", banner)
		return ErrIncorrectData{msg: fmt.Sprintf("данные %s некорректны в данном поле должна быть ссылка", banner.Content.Url)}
	}

	return nil

}
