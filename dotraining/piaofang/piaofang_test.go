package piaofang

import (
	"fmt"
	"github.com/lzhphantom/go-spider/url"
	"testing"
)

func TestGetBoxOfficeInfo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url.BoxOfficeURL},
	}
	for _, test := range tests {
		fmt.Println(GetBoxOfficeInfo(test.url))
	}
}
