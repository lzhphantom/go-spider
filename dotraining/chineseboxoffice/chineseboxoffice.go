package chineseboxoffice

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/lzhphantom/go-spider/downloader"
	"github.com/lzhphantom/go-spider/inital"
	"github.com/lzhphantom/go-spider/model"
	"github.com/lzhphantom/go-spider/url"
	"github.com/tidwall/gjson"
	"strings"
	"time"
	"unsafe"
)

type RealTimeMovie struct {
	Rank      string `json:"rank"`
	MovieName string `json:"movie_name"`
	BoxOffice string `json:"box_office"`

	Rate         string `json:"rate"`
	SumBoxOffice string `json:"sum_box_office"`
	TicketRate   string `json:"ticket_rate"`
	Day          string `json:"day"`
}

var areaMap = make(map[string]string)
var urlList []string
var areaList []string
var total int64

func init() {
	areaMap = map[string]string{
		"中国":   "50",
		"中国香港": "37",
		"中国台湾": "40",
		"美国":   "1",
		"英国":   "25",
		"德国":   "16",
		"法国":   "4",
		"日本":   "30",
		"加拿大":  "2",
		"意大利":  "7",
	}

	for key, value := range areaMap {
		var oneUrl string
		oneUrl = fmt.Sprintf(url.Cbooo_RealUrl, value, 1)
		urlList = append(urlList, oneUrl)
		areaList = append(areaList, key)
	}
}

func getCboooRankList() {
	inital.DBInit()
	for indexArea, oneUrl := range urlList {
		time.Sleep(1 * time.Millisecond * 1000 * 10)
		response, err := downloader.GetHttpResponse(oneUrl, false)
		if err != nil {
			return
		}
		responseString := (*string)(unsafe.Pointer(&response))
		totalPage := gjson.Parse(*responseString).Get("tPage").Int()
		total += gjson.Parse(*responseString).Get("tCount").Int()
		fmt.Println("totalPage", totalPage, "current_count", gjson.Parse(*responseString).Get("tCount").Int(), "total:", total)
		fmt.Println(oneUrl)

		for index := 1; index <= int(totalPage); index++ {
			newOneUrl := strings.Replace(oneUrl, "pIdex=1", fmt.Sprintf("pIndex=%d", index), -1)
			fmt.Println(newOneUrl)
			go func(url string) {
				time.Sleep(1 * time.Millisecond * 1000 * 10)
				newResponse, err := downloader.GetHttpResponse(newOneUrl, false)
				if err != nil {
					return
				}
				newResponseString := (*string)(unsafe.Pointer(&newResponse))
				info := gjson.Parse(*newResponseString).Get("pData")
				for _, oneInfo := range info.Array() {
					var oneRank model.ChoooRankListInfo
					oneRank = model.ChoooRankListInfo{
						Ranking:      oneInfo.Get("Ranking").String(),
						MovieID:      oneInfo.Get("ID").String(),
						MovieEnName:  oneInfo.Get("MovieEnName").String(),
						MovieName:    oneInfo.Get("MovieName").String(),
						ReleaseYear:  oneInfo.Get("releaseYear").String(),
						DefaultImage: oneInfo.Get("defaultImage").String(),
						BoxOffice:    oneInfo.Get("BoxOffice").String(),
						Country:      areaList[indexArea],
					}
					fmt.Println(oneRank)
					inital.DataBase.Create(&oneRank)
				}
			}(newOneUrl)
		}
	}
	fmt.Println("total", total)
}

//中国票房首页 --- 票房榜信息
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

// 中国票房 --- 影库
func getAreaList(url string) {
	response, err := downloader.GetHttpResponse(url, false)
	if err != nil {
		return
	}

	stringReader := (*string)(unsafe.Pointer(&response))
	document, err := goquery.NewDocumentFromReader(strings.NewReader(*stringReader))

	if err != nil {
		return
	}
	document.Find("#selArea option").Each(func(i int, selection *goquery.Selection) {
		country := selection.Text()
		value, _ := selection.Attr("value")
		fmt.Println(country, value)
	})

	document.Find("#selType option").Each(func(i int, selection *goquery.Selection) {
		movieType := selection.Text()
		value, _ := selection.Attr("value")
		fmt.Println(movieType, value)
	})

	document.Find("#selYear option").Each(func(i int, selection *goquery.Selection) {
		year := selection.Text()
		value, _ := selection.Attr("value")
		fmt.Println(year, value)
	})
}
