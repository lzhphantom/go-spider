package chineseboxoffice

import (
	"github.com/lzhphantom/go-spider/url"
	"testing"
)

func TestGetRealTimeList(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url.ChineseBoxOffice},
	}

	for _, test := range tests {
		getRealTimeList(test.url)
	}
}

func TestGetAreaList(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url.CBooo_Movie},
	}

	for _, test := range tests {
		getAreaList(test.url)
	}
}
