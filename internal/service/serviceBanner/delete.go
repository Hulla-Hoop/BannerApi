package servicebanner

import (
	"errors"
	"regexp"
	"strconv"
)

func (c *serviceBanner) Delete(reqId string, id string) error {

	idi, err := c.idCheckAndConvert(reqId, id)
	if err != nil {
		return err
	}
	err = c.db.Delete(reqId, idi)
	if err != nil {
		return err
	}

	return nil

}

// проверяет ID на наличие некорректных даннных и в случае их отсутствия переводит его в int
func (c *serviceBanner) idCheckAndConvert(reqId string, id string) (int, error) {
	pattern := `^[0-9]*$`
	match, _ := regexp.MatchString(pattern, id)
	if !match {
		c.logger.WithField("carCatalog.idCheckAndConvert", reqId).Error("некорректные данные ", id)
		return 0, errors.New("некорректные данные")
	}
	idi, err := strconv.Atoi(id)
	if err != nil {
		c.logger.WithField("carCatalog.idCheckAndConvert", reqId).Error("Ннеудалось преобразовать ", id)
		return 0, err
	}

	return idi, nil
}
