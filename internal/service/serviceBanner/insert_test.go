package servicebanner

import (
	"banner/internal/config"
	"banner/internal/logger"
	"banner/internal/model"
	"banner/internal/repo"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_serviceBanner_validate(t *testing.T) {
	type fields struct {
		logger *logrus.Logger
		db     repo.Repos
		cfg    *config.ConfigRemoteApi
	}
	type args struct {
		reqId  string
		banner model.BannerHttp
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "некорректный url",
			fields: fields{
				logger: logger.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId: "Test",
				banner: model.BannerHttp{
					Banner_id:  1,
					Tags_id:    model.Tags{1, 2, 3},
					Feature_id: 1,
					Content: model.Banner{
						Title: "Test",
						Text:  "Test",
						Url:   "Test",
					},
					Is_active: true,
				},
			},

			wantErr: true,
		},
		{
			name: "некорректный feature_id",
			fields: fields{
				logger: logger.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId: "Test",
				banner: model.BannerHttp{
					Banner_id:  1,
					Tags_id:    model.Tags{1, 2, 3},
					Feature_id: -1,
					Content: model.Banner{
						Title: "Test",
						Text:  "Test",
						Url:   "http://Test",
					},
					Is_active: true,
				},
			},
			wantErr: true,
		},
		{
			name: "некорректный tag_id",
			fields: fields{
				logger: logger.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId: "Test",
				banner: model.BannerHttp{
					Banner_id:  1,
					Tags_id:    model.Tags{1, -2, 3},
					Feature_id: 1,
					Content: model.Banner{
						Title: "Test",
						Text:  "Test",
						Url:   "ftp://Test",
					},
					Is_active: true,
				},
			},
			wantErr: false,
		},
		{
			name: "корректные данные",
			fields: fields{
				logger: logger.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId: "Test",
				banner: model.BannerHttp{
					Banner_id:  1,
					Tags_id:    model.Tags{1, 2, 3},
					Feature_id: 1,
					Content: model.Banner{
						Title: "Test",
						Text:  "Test",
						Url:   "https://Test",
					},
					Is_active: true,
				},
			},
			wantErr: false,
		},
		{
			name: "все теги некорректные",
			fields: fields{
				logger: logger.New(),
				db:     nil,
				cfg:    nil,
			},
			args: args{
				reqId: "Test",
				banner: model.BannerHttp{
					Banner_id:  1,
					Tags_id:    model.Tags{-1, -2, -3},
					Feature_id: 1,
					Content: model.Banner{
						Title: "Test",
						Text:  "Test",
						Url:   "https://Test",
					},
					Is_active: true,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &serviceBanner{
				logger: tt.fields.logger,
				db:     tt.fields.db,
				cfg:    tt.fields.cfg,
			}
			err := c.validate(tt.args.reqId, tt.args.banner)
			if (err != nil) != tt.wantErr {
				t.Errorf("serviceBanner.validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
