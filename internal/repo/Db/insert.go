package psql

import (
	"banner/internal/model"
	"fmt"
)

func (p *psql) Insert(reqId string, b model.BannerDB, t model.Tags) (int, error) {

	var id int

	p.logger.WithField("psql.Insert", reqId).Debug("Полученные данные", b, t)

	query := fmt.Sprintf(`INSERT INTO banner ( title, text, url, active,created_at, updated_at) 
	VALUES ( '%s', '%s', '%s', '%t','%s','%s') 
	returning id;`,
		b.Title, b.Text, b.Url, b.Active, b.Created_at, b.Updated_at)

	p.logger.WithField("psql.Insert", reqId).Debug("Запрос на вставку", query)

	err := p.dB.QueryRow(query).Scan(&id)
	if err != nil {
		p.logger.WithField("psql.Insert", reqId).Error(err)
		return 0, err
	}

	for _, v := range t {
		err := p.dB.QueryRow("INSERT INTO chains (banner_id, tags_id, feature_id) VALUES ($1, $2, $3)", id, v, b.Feature).Err()
		if err != nil {
			p.logger.WithField("psql.Insert", reqId).Error(err)
			return 0, err
		}
	}

	return id, nil
}
