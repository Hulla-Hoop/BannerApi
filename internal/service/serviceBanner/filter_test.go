package servicebanner

import (
	"banner/internal/repo"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_serviceBanner_filterValidate(t *testing.T) {
	type fields struct {
		logger *logrus.Logger
		db     repo.Repos
	}
	type args struct {
		reqId     string
		tegID     string
		featureID string
		limit     string
		offset    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		{
			name: "все корректо",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
			},
			args: args{
				reqId:     "test",
				limit:     "1",
				offset:    "1",
				tegID:     "1",
				featureID: "1",
			},
			want: map[string]string{"limit": "1", "offset": "1", "feature": "1", "tag": "1"},
		},
		{
			name: "все некорректно",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
			},
			args: args{
				reqId:     "1",
				limit:     "-2",
				offset:    "-5",
				tegID:     "-2",
				featureID: "-4",
			},
			want: map[string]string{},
		},
		{
			name: "только лимит и офсет",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
			},
			args: args{
				reqId:     "1",
				limit:     "1",
				offset:    "1",
				featureID: "ssd",
				tegID:     "dasd",
			},
			want: map[string]string{"limit": "1", "offset": "1"},
		},
		{
			name: "только фича и тег",
			fields: fields{
				logger: logrus.New(),
				db:     nil,
			},
			args: args{
				reqId:     "1",
				limit:     "-100",
				offset:    "-90",
				tegID:     "1",
				featureID: "1",
			},
			want: map[string]string{"tag": "1", "feature": "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &serviceBanner{
				logger: tt.fields.logger,
				db:     tt.fields.db,
			}
			if got := c.filterValidate(tt.args.reqId, tt.args.tegID, tt.args.featureID, tt.args.limit, tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("serviceBanner.filterValidate() = %v, want %v", got, tt.want)
			}
		})
	}
}
