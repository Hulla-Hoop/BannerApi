package model

import "time"

type Banner struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

type Tags []int

type BannerHttp struct {
	Banner_id  int    `json:"banner_id"`
	Tags_id    Tags   `json:"tag_ids"`
	Feature_id int    `json:"feature_id"`
	Content    Banner `json:"content"`
	Is_active  bool   `json:"is_active"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

func (b *BannerHttp) BFTOTagsAndBannerDB() (BannerDB, Tags) {

	return BannerDB{b.Banner_id, b.Feature_id, b.Content.Title, b.Content.Text, b.Content.Url, b.Is_active, time.Now().Format(time.DateTime), time.Now().Format(time.DateTime)}, b.Tags_id

}

type BannerDB struct {
	Id         int
	Feature    int
	Title      string
	Text       string
	Url        string
	Active     bool
	Created_at string
	Updated_at string
}

func (b BannerDB) TOTagsAndBannerFilter(Tags) BannerHttp {
	return BannerHttp{b.Id, Tags{}, b.Feature, Banner{b.Title, b.Text, b.Url}, b.Active, b.Created_at, b.Updated_at}
}
