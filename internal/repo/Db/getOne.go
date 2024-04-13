package psql

import (
	"banner/internal/model"
	"fmt"
)

func (p *psql) GetOne(reqId string, tag_id int, feature_id int, last bool, root bool) (model.Banner, error) {
	p.logger.WithField("psql.GetOne", reqId).Debug("Полученные данные -- ", tag_id, feature_id, last)

	var b model.Banner

	ok, err := p.chekRowForGetOne(reqId, feature_id, tag_id)
	if err != nil {
		p.logger.WithField("psql.GetOne", reqId).Error(err)
		return model.Banner{}, err
	}

	if !ok {
		return model.Banner{}, ErrNotFound{msg: fmt.Sprintf("Баннер с feature_id %d и tag_id %d не найден", feature_id, tag_id)}
	}

	if !root {
		err = p.dB.QueryRow(`SELECT banner.title,banner.text,banner.url FROM banner
	JOIN chains
	ON banner.id=chains.banner_id
	WHERE chains.feature_id=$1 AND chains.tags_id=$2 AND banner.active=true;
	`, feature_id, tag_id).Scan(&b.Title, &b.Text, &b.Url)
	} else {
		err = p.dB.QueryRow(`SELECT banner.title,banner.text,banner.url FROM banner
	JOIN chains
	ON banner.id=chains.banner_id
	WHERE chains.feature_id=$1 AND chains.tags_id=$2;
	`, feature_id, tag_id).Scan(&b.Title, &b.Text, &b.Url)
	}

	if err != nil {

		p.logger.WithField("psql.GetOne", reqId).Error(err)

		return model.Banner{}, err
	}

	return b, nil

}

func (p *psql) chekRowForGetOne(reqId string, feature_id int, tag_id int) (bool, error) {

	var ok bool

	p.logger.WithField("psql.chekRowForGetOne", reqId).Debug("Полученные данные -- ", feature_id, "---", tag_id)

	err := p.dB.QueryRow("SELECT EXISTS(SELECT * FROM chains WHERE feature_id = $1, tags_id = $2)", feature_id, tag_id).Scan(&ok)

	if err != nil {
		p.logger.WithField("psql.chekRow", reqId).Error(err)
		return false, err
	}

	return ok, nil

}
