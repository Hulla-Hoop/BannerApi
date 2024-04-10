package psql

import "fmt"

func (p *psql) Delete(reqId string, id int) error {

	p.logger.WithField("psql.Delete", reqId).Debug("Полученные данные", id)

	ok, err := p.chekRowForDelete(reqId, id)
	if err != nil {
		p.logger.WithField("psql.Delete", reqId).Error(err)
		return err
	}
	if !ok {
		return ErrNotFound{msg: fmt.Sprintf("Баннер с id %d не найден", id)}
	}

	err = p.dB.QueryRow("DELETE FROM banner WHERE id = $1", id).Err()
	if err != nil {
		p.logger.WithField("psql.Delete", reqId).Error(err)
		return err
	}

	err = p.dB.QueryRow("DELETE FROM chains WHERE banner_id = $1", id).Err()
	if err != nil {
		p.logger.WithField("psql.Delete", reqId).Error(err)
		return err
	}

	return nil
}

func (p *psql) chekRowForDelete(reqId string, id int) (bool, error) {

	var ok bool
	p.logger.WithField("psql.chekRow", reqId).Debug("Полученные данные", id)
	err := p.dB.QueryRow("SELECT EXISTS(SELECT * FROM banner WHERE id = $1)", id).Scan(&ok)
	if err != nil {
		p.logger.WithField("psql.chekRow", reqId).Error(err)
		return false, err
	}
	return ok, nil
}
