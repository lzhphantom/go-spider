package maoyan

import (
	"fmt"
	"github.com/lzhphantom/go-spider/url"
	"testing"
)

func TestGetMaoYanMovieInfo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url.GetMaoYanUrl(2019)},
	}

	for _, test := range tests {
		fmt.Println(GetMaoYanMovieInfo(test.url))
	}

}
