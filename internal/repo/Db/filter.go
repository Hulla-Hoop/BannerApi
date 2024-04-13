package psql

import (
	"banner/internal/model"
	"database/sql"
	"fmt"
)

func (p *psql) Filter(reqId string, filter map[string]string) ([]model.BannerHttp, error) {

	query := p.queryFilter(filter)

	p.logger.WithField("psql.Filter", reqId).Debug("Запрос на фильтрацию:  ", query)

	rows, err := p.dB.Query(query)
	if err != nil {
		p.logger.WithField("psql.Filter", reqId).Error(err)
		return nil, err
	}

	defer rows.Close()

	var bannerSL []model.BannerHttp

	for rows.Next() {
		var bannerDB model.BannerDB

		err := rows.Scan(&bannerDB.Id, &bannerDB.Title, &bannerDB.Text, &bannerDB.Url, &bannerDB.Active, &bannerDB.Created_at, &bannerDB.Updated_at)

		if err != nil {
			if err == sql.ErrNoRows {
				p.logger.WithField("psql.Filter", reqId).Debug("Баннеры не найдены")
				continue
			}
			p.logger.WithField("psql.Filter", reqId).Error(err)

			return nil, err

		}

		err = p.dB.QueryRow("Select feature_id from chains where banner_id = $1", bannerDB.Id).Scan(&bannerDB.Feature)
		if err != nil {

			if err == sql.ErrNoRows {
				p.logger.WithField("psql.Filter", reqId).Debug("Баннеры не найдены")
				continue
			}
			p.logger.WithField("psql.Filter", reqId).Error(err)

			return nil, err
		}

		var tegs model.Tags

		row, err := p.dB.Query("Select tags_id from chains where banner_id = $1", bannerDB.Id)
		if err != nil {

			if err == sql.ErrNoRows {
				p.logger.WithField("psql.Filter", reqId).Debug("Баннеры не найдены")
				continue
			}
			p.logger.WithField("psql.Filter", reqId).Error(err)

			return nil, err
		}

		defer rows.Close()

		for row.Next() {
			var tegID int
			err := row.Scan(&tegID)
			if err != nil {

				if err == sql.ErrNoRows {
					p.logger.WithField("psql.Filter", reqId).Debug("Баннеры не найдены")
					continue
				}
				p.logger.WithField("psql.Filter", reqId).Error(err)

				return nil, err
			}
			tegs = append(tegs, tegID)
		}

		b := bannerDB.TOTagsAndBannerFilter(tegs)

		p.logger.WithField("psql.Filter", reqId).Debug("Полученные данные", b)

		bannerSL = append(bannerSL, b)
	}

	return bannerSL, nil
}

// генирирует строку запроса на основе полученной мапы которая содержит параметры фильтрации
func (p *psql) queryFilter(filter map[string]string) string {

	p.logger.WithField("psql.queryFilter", "").Debug("Полученные данные", filter)

	query := `SELECT *
	FROM banner `

	if len(filter) == 0 {

		query += "WHERE active = true "
		return query

	} else {

		query += "WHERE active = true "

		offset, ok := filter["offset"]
		limit, es := filter["limit"]

		if ok && es {
			query += fmt.Sprintf("AND id > %s", offset)

			field, ok := filter["feature"]
			if ok {
				query += fmt.Sprintf(" AND id IN (SELECT banner_id FROM chains WHERE feature_id = %s) ", field)
			}

			value, es := filter["tag"]
			if es {
				query += fmt.Sprintf(" AND id IN (SELECT banner_id FROM chains WHERE tags_id = %s) ", value)
			}

			query += fmt.Sprintf("LIMIT %s", limit)

			return query
		} else {

			field, ok := filter["feature"]
			if ok {
				query += fmt.Sprintf(" AND id IN (SELECT banner_id FROM chains WHERE feature_id = %s) ", field)
			}

			value, es := filter["tag"]
			if es {
				query += fmt.Sprintf(" AND id IN (SELECT banner_id FROM chains WHERE tags_id = %s) ", value)
			}

			return query
		}

	}

}
