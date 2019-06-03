package piaofang

import (
	"fmt"
	"github.com/lzhphantom/go-spider/downloader"
	"regexp"
	"unsafe"
)

type MovieInfo struct {
	Rank            string `json:"rank"`
	MoveName        string `json:"movie_name"`
	Publish         string `json:"publish"`
	MovieType       string `json:"movie_type"`
	Director        string `json:"director"`
	GlobalBoxOffice string `json:"global_box_office"`
}

var (
	rankPatter            = `<td class="num">(.*?)</td>`
	movieNamePatter       = `<td class="title">(.*?)<span>(.*?)</span></td>`
	PublishPatter         = `<td class="year">(.*?)</td>`
	MovieTypePatter       = `<td class="type">(.*?)</td>`
	directorPatter        = `<td class="daoyan">(.*?)</td>`
	globalBoxOfficePatter = `<td class="piaofang"><span>(.*?)</span>(.*?)</td>`
)

func GetBoxOfficeInfo(url string) ([]MovieInfo, error) {
	responseByte, err := downloader.GetHttpResponse(url, true)
	if err != nil {
		return nil, err
	}

	responseString := (*string)(unsafe.Pointer(&responseByte))
	var rank []string
	reRank := regexp.MustCompile(rankPatter)
	for index, subMatch := range reRank.FindAllStringSubmatch(*responseString, -1) {
		if len(subMatch[1]) > 34 {
			subMatch[1] = subMatch[1][34:len(subMatch[1])]
		}
		fmt.Println(index, subMatch[1])
		rank = append(rank, subMatch[1])
	}

	var movieNameList []string
	movieName := regexp.MustCompile(movieNamePatter)
	for index, subMatch := range movieName.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, subMatch[1], subMatch[2])
		movieNameList = append(movieNameList, subMatch[1])
	}

	var publishList []string
	publish := regexp.MustCompile(PublishPatter)
	for index, subMatch := range publish.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, subMatch[1])
		publishList = append(publishList, subMatch[1])
	}

	var movieTypeList []string
	movieType := regexp.MustCompile(MovieTypePatter)
	for index, subMatch := range movieType.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, subMatch[1])
		movieTypeList = append(movieTypeList, subMatch[1])
	}

	var directorList []string
	director := regexp.MustCompile(directorPatter)
	for index, subMatch := range director.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, subMatch[1])
		directorList = append(directorList, subMatch[1])
	}

	var globalBoxOfficeList []string
	globalBoxOffice := regexp.MustCompile(globalBoxOfficePatter)
	for index, subMatch := range globalBoxOffice.FindAllStringSubmatch(*responseString, -1) {
		fmt.Println(index, subMatch[1])
		globalBoxOfficeList = append(globalBoxOfficeList, subMatch[1])
	}

	var result []MovieInfo
	for i := 0; i < len(rank); i++ {
		var movieInfo MovieInfo
		movieInfo = MovieInfo{
			rank[i],
			movieNameList[i],
			publishList[i],
			movieTypeList[i],
			directorList[i],
			globalBoxOfficeList[i],
		}
		result = append(result, movieInfo)
	}
	return result, nil
}
