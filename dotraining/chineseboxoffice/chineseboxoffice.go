package chineseboxoffice

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/lzhphantom/go-spider/downloader"
	"strings"
	"unsafe"
)

type RealTimeMovie struct {
	Rank         string `json:"rank"`
	MovieName    string `json:"movie_name"`
	BoxOffice    string `json:"box_office"`
	Rate         string `json:"rate"`
	SumBoxOffice string `json:"sum_box_office"`
	TicketRate   string `json:"ticket_rate"`
	Day          string `json:"day"`
}

func getRealTimeList(url string) {
	response, err := downloader.GetHttpResponse(url, false)
	if err != nil {
		return
	}
	stringReader := (*string)(unsafe.Pointer(&response))
	document, err := goquery.NewDocumentFromReader(strings.NewReader(*stringReader))

	if err != nil {
		return
	}
	document.Find("#topdatatr tr").Each(func(i int, selection *goquery.Selection) {
		var RealTimeTicket RealTimeMovie
		RealTimeTicket = RealTimeMovie{
			selection.Find("td").Eq(0).Text(),
			selection.Find("td").Eq(1).Text(),
			selection.Find("td").Eq(2).Text(),
			selection.Find("td").Eq(3).Text(),
			selection.Find("td").Eq(4).Text(),
			selection.Find("td").Eq(5).Text(),
			selection.Find("td").Eq(6).Text(),
		}
		fmt.Println(RealTimeTicket)
	})
}
