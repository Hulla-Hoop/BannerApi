package repo

import "banner/internal/model"

type Repos interface {
	Insert(reqId string, b model.BannerDB, t model.Tags) (int, error)
	Update(reqId string, b model.BannerDB, t model.Tags) error
	Filter(reqId string, filter map[string]string) ([]model.BannerDB, error)
	GetOne(reqId string, tag_id int, feature_id int, last bool) (model.Banner, error)
	Delete(reqId string, id int) error
}
