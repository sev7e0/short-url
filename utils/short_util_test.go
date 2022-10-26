package utils

import "testing"

func TestGenShortUrl(t *testing.T) {
	type args struct {
		oriUrl string
	}
	tests := []struct {
		name         string
		args         args
		wantShortUrl string
	}{
		{"1", args{oriUrl: "www.didi.com"}, "4zybfT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotShortUrl := GenShortUrl(tt.args.oriUrl); gotShortUrl != tt.wantShortUrl {
				t.Errorf("GenShortUrl() = %v, want %v", gotShortUrl, tt.wantShortUrl)
			}
		})
	}
}
