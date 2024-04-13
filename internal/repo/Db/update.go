package psql

import (
	"banner/internal/model"
	"fmt"
)

func (p *psql) Update(reqId string, b model.BannerDB, t model.Tags) error {
	p.logger.WithField("psql.Update", reqId).Debug("Полученные данные banner -- ", b, " tags -- ", t)

	err := p.chekRowOnID(reqId, b.Id)
	if err != nil {
		return err
	}

	query := p.getQuery(reqId, b)

	err = p.dB.QueryRow(query).Err()
	if err != nil {
		p.logger.WithField("psql.Update", reqId).Error(err)
		return err
	}
	// удаляем старые теги
	err = p.dB.QueryRow("DELETE FROM chains WHERE banner_id = $1", b.Id).Err()
	if err != nil {
		p.logger.WithField("psql.Update", reqId).Error(err)
		return err
	}

	// добавляем новые теги
	for _, v := range t {
		err := p.dB.QueryRow("INSERT INTO chains (banner_id, tags_id, feature_id) VALUES ($1, $2, $3)", b.Id, v, b.Feature).Err()
		if err != nil {
			p.logger.WithField("psql.Update", reqId).Error(err)
			return err
		}
	}

	return nil
}

// генерируем строку запроса на основе не пустых полей структуры Banner
func (p *psql) getQuery(reqId string, b model.BannerDB) string {

	query := `
	UPDATE banner 
	SET `

	if b.Title != "" {
		query += fmt.Sprintf("title = '%s',", b.Title)
	}
	if b.Text != "" {
		query += fmt.Sprintf("text = '%s',", b.Text)
	}
	if b.Url != "" {
		query += fmt.Sprintf("url = '%s',", b.Url)
	}
	if !b.Active {
		query += fmt.Sprintf("active = '%t',", b.Active)
	}

	query = query + fmt.Sprintf("updated_at = '%s' WHERE id = %d;", b.Updated_at, b.Id)

	p.logger.WithField("psql.Update", reqId).Debug("Запрос на обновление", query)

	return query
}
