package golibrary

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetFaviconHash(t *testing.T) {
	type args struct {
		url string
		r   *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{name: "1", args: args{url: "https://www.baidu.com/favicon.ico"}, want: -1588080585, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.r, _ = http.Get(tt.args.url)
			got, _ := GetFaviconHash(tt.args.r)
			assert.Equalf(t, tt.want, got, "GetFaviconHash(%v)", tt.args.r)
		})
	}
}
