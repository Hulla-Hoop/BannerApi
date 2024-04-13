package service

import "banner/internal/model"

type ServiceBanner interface {
	Filter(reqId string, tegID string, featureID string, limit string, offset string) ([]model.BannerHttp, error)
	GetOne(reqId string, tegID string, featureID string, las string, role bool) (model.Banner, error)
	Insert(reqId string, body []byte) (int, error)
	Update(reqId string, id string, ban []byte) error
	Delete(reqId string, id string) error
}
