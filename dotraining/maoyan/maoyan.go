package maoyan

import (
	"fmt"
	"github.com/lzhphantom/go-spider/downloader"
	"regexp"
	"unsafe"
)

type MovieMaoYanInfo struct {
	Rank              string `json:"rank"`
	MovieName         string `json:"movie_name"`
	StartYear         string `json:"start_year"`
	BoxOfficeIncome   string `json:"box_office_income"`
	AvgEachTicket     string `json:"avg_each_ticket"`
	PerformancePeople string `json:"performance_people"`
	Url               string `json:"url"`
}

var (
	maoyanRootUrl               = `https://piaofang.maoyan.com/`
	rankPatter                  = `<li class="col0">(.*?)</li>`
	movieNameAndStartYearPatter = `<li class="col1">
        <p class="first-line">(.*?)</p>
        <p class="second-line">(.*?) 上映</p>
    </li>`
	boxOfficeIncomePatter    = `<li class="col2 tr">(.*?)</li>`
	avgEachTicketPatter      = `<li class="col3 tr">(.*?)</li>`
	performancePeoplePattern = `<li class="col4 tr">(.*?)</li>`
	urlPattern               = `<ul class="row" data-com="hrefTo,href:'(.*?)'" data-loaded="true">`
)

func GetMaoYanMovieInfo(url string) ([]MovieMaoYanInfo, error) {
	response, err := downloader.GetHttpResponse(url, false)
	if err != nil {
		return nil, err
	}
	responseString := (*string)(unsafe.Pointer(&response))
	//fmt.Println(*responseString)

	var urlList []string
	urlRegexp := regexp.MustCompile(urlPattern)
	for index, one := range urlRegexp.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, one[1])
		urlList = append(urlList, one[1])
	}

	var rankList []string
	rankRegexp := regexp.MustCompile(rankPatter)
	for index, rank := range rankRegexp.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, rank[1])
		if index > 0 {
			rankList = append(rankList, rank[1])
		}
	}

	var movieNameList []string
	var startYearList []string
	movieNameAndStartYearRegexp := regexp.MustCompile(movieNameAndStartYearPatter)
	for index, movie := range movieNameAndStartYearRegexp.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, movie[1], movie[2])
		movieNameList = append(movieNameList, movie[1])
		startYearList = append(startYearList, movie[2])
	}

	var boxOfficeIncomeList []string
	boxOfficeIncomeRegexp := regexp.MustCompile(boxOfficeIncomePatter)
	for index, boxOffice := range boxOfficeIncomeRegexp.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, boxOffice[1])
		boxOfficeIncomeList = append(boxOfficeIncomeList, boxOffice[1])
	}

	var avgEachTicketList []string
	avgEachTicketRegexp := regexp.MustCompile(avgEachTicketPatter)
	for index, avgeach := range avgEachTicketRegexp.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, avgeach[1])
		avgEachTicketList = append(avgEachTicketList, avgeach[1])
	}

	var performancePeopleList []string
	performancePeopleRegexp := regexp.MustCompile(performancePeoplePattern)
	for index, performance := range performancePeopleRegexp.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, performance[1])
		performancePeopleList = append(performancePeopleList, performance[1])
	}

	var result []MovieMaoYanInfo
	for index := 0; index < len(movieNameList); index++ {
		var myInfo MovieMaoYanInfo
		myInfo = MovieMaoYanInfo{
			Rank:      rankList[index],
			MovieName: movieNameList[index],
			StartYear: startYearList[index],
			Url:       urlList[index],
		}
		result = append(result, myInfo)
	}
	return result, nil
}
