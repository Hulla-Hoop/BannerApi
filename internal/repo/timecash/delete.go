package timecash

func (c *timeCash) Delete(reqId string, id int) error {

	return c.db.Delete(reqId, id)
}
