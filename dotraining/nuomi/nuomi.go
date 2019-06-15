package nuomi

import (
	"fmt"
	"github.com/lzhphantom/go-spider/downloader"
	"github.com/tidwall/gjson"
	"unsafe"
)

type CommonRankListFiled struct {
	MovieId   string `json:"movie_id"`
	MovieName string `json:"movie_name"`
	Index     string `json:"index"`
}

type BoxRankListNuoMiInfo struct {
	CommonRankListFiled
	RealTimeBox string `json:"real_time_box"`
}

type FocusRankListNuoMiInfo struct {
	CommonRankListFiled
	MovieType string `json:"movie_type"`
	FocusNum  string `json:"focus_num"`
}

type SearchRankListNuoMiInfo struct {
	CommonRankListFiled
	MovieType   string `json:"movie_type"`
	SearchCount string `json:"search_count"`
}

type ReputationRankList struct {
	CommonRankListFiled
	MovieType  string `json:"movie_type"`
	Reputation string `json:"reputation"`
}

type AgeRankList struct {
	CommonRankListFiled
	MovieType string `json:"movie_type"`
}

type AnnualReputationMovieRankList struct {
	AgeRankList
}

type RegionRankList struct {
	AgeRankList
}

func GetRankNuoMiInfo(url string) {
	response, err := downloader.GetHttpResponse(url, false)
	if err != nil {
		return
	}
	responseString := (*string)(unsafe.Pointer(&response))
	//box := gjson.Parse(string([]byte(response))).Get("boxRankList")
	box := gjson.Parse(*responseString).Get("boxRankList")
	var boxes []BoxRankListNuoMiInfo
	for _, oneBoxRank := range box.Array() {
		var box BoxRankListNuoMiInfo
		box = BoxRankListNuoMiInfo{RealTimeBox: oneBoxRank.Get("realTimeBox").String(),
			CommonRankListFiled: CommonRankListFiled{
				oneBoxRank.Get("movieId").String(),
				oneBoxRank.Get("movieName").String(),
				oneBoxRank.Get("index").String(),
			}}
		boxes = append(boxes, box)
	}

	focus := gjson.Parse(*responseString).Get("focuRankList")
	var focusRankList []FocusRankListNuoMiInfo
	for _, OneFocus := range focus.Array() {
		var oneFocusRank = FocusRankListNuoMiInfo{
			CommonRankListFiled{
				OneFocus.Get("movieId").String(),
				OneFocus.Get("movieName").String(),
				OneFocus.Get("index").String(),
			},
			OneFocus.Get("movieType").String(),
			OneFocus.Get("focusNum").String(),
		}
		focusRankList = append(focusRankList, oneFocusRank)
	}

	search := gjson.Parse(*responseString).Get("searchRankList")
	var searchRankList []SearchRankListNuoMiInfo
	for _, OneSearch := range search.Array() {
		var search = SearchRankListNuoMiInfo{
			CommonRankListFiled: CommonRankListFiled{
				OneSearch.Get("movieId").String(),
				OneSearch.Get("movieName").String(),
				OneSearch.Get("index").String(),
			},
			SearchCount: OneSearch.Get("searchCount").String(),
		}
		searchRankList = append(searchRankList, search)
	}

	reputation := gjson.Parse(*responseString).Get("reputationRankList")
	var reputataionList []ReputationRankList

	for _, one := range reputation.Array() {
		var ore = ReputationRankList{
			CommonRankListFiled: CommonRankListFiled{
				one.Get("movieId").String(),
				one.Get("movieName").String(),
				one.Get("index").String(),
			},
			Reputation: one.Get("reputation").String(),
		}
		reputataionList = append(reputataionList, ore)
	}
	ageRankList := gjson.Parse(string([]byte(response))).Get("ageRankList")
	var ageRankLists []AgeRankList
	for _, OneAgeRank := range ageRankList.Array() {
		var oneAgeRankList AgeRankList
		oneAgeRankList = AgeRankList{
			CommonRankListFiled: CommonRankListFiled{
				MovieId:   OneAgeRank.Get("movieId").String(),
				MovieName: OneAgeRank.Get("movieName").String(),
				Index:     OneAgeRank.Get("index").String(),
			},
			MovieType: OneAgeRank.Get("movieType").String(),
		}
		ageRankLists = append(ageRankLists, oneAgeRankList)
	}

	annualBoxMovie := gjson.Parse(string([]byte(response))).Get("annualBoxMovieRankList")

	var annualBoxMovies []AnnualReputationMovieRankList
	for _, OneAnnualBox := range annualBoxMovie.Array() {
		var OneAnnualRank AnnualReputationMovieRankList
		OneAnnualRank = AnnualReputationMovieRankList{
			AgeRankList: AgeRankList{
				CommonRankListFiled: CommonRankListFiled{
					MovieId:   OneAnnualBox.Get("movieId").String(),
					MovieName: OneAnnualBox.Get("movieName").String(),
					Index:     OneAnnualBox.Get("index").String(),
				},
				MovieType: OneAnnualBox.Get("movieType").String(),
			},
		}
		annualBoxMovies = append(annualBoxMovies, OneAnnualRank)

	}
	annualReputation := gjson.Parse(string([]byte(response))).Get("annualReputationMovieRankList")
	var annualReputations []AnnualReputationMovieRankList
	for _, OneAnnual := range annualReputation.Array() {
		var OneAnnualRank AnnualReputationMovieRankList
		OneAnnualRank = AnnualReputationMovieRankList{
			AgeRankList: AgeRankList{
				CommonRankListFiled: CommonRankListFiled{
					MovieId:   OneAnnual.Get("movieId").String(),
					MovieName: OneAnnual.Get("movieName").String(),
					Index:     OneAnnual.Get("index").String(),
				},
				MovieType: OneAnnual.Get("movieType").String(),
			},
		}
		annualReputations = append(annualReputations, OneAnnualRank)
	}
	regionRank := gjson.Parse(string([]byte(response))).Get("regionRankList")

	var regionRanks []RegionRankList
	for _, OneRegion := range regionRank.Array() {
		var OneRank RegionRankList
		OneRank = RegionRankList{
			AgeRankList: AgeRankList{
				CommonRankListFiled: CommonRankListFiled{
					MovieId:   OneRegion.Get("movieId").String(),
					MovieName: OneRegion.Get("movieName").String(),
					Index:     OneRegion.Get("index").String(),
				},
				MovieType: OneRegion.Get("movieType").String(),
			},
		}
		regionRanks = append(regionRanks, OneRank)

	}

	fmt.Println(boxes, focusRankList, searchRankList, reputataionList, ageRankLists, annualBoxMovies,
		annualReputations, regionRanks)
}
