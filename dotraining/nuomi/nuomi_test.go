package nuomi

import (
	"github.com/lzhphantom/go-spider/url"
	"testing"
)

func TestGetRankNuoMiInfo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url.NuoMiRankURL},
	}

	for _, test := range tests {
		GetRankNuoMiInfo(test.url)
	}
}
