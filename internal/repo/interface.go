package repo

import "banner/internal/model"

type Repos interface {
	Insert(reqId string, b model.BannerDB, t model.Tags) (int, error)
	Update(reqId string, b model.BannerDB, t model.Tags) error
	Filter(reqId string, filter map[string]string) ([]model.BannerHttp, error)
	GetOne(reqId string, tag_id int, feature_id int, last bool, role bool) (model.Banner, error)
	Delete(reqId string, id int) error
}

//TODO: сделать слой кеширования
