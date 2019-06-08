package chineseFilm

import (
	"fmt"
	"github.com/lzhphantom/go-spider/downloader"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"strconv"
	"unsafe"
)

var body = make(map[string]string)

func getBody() map[string]string {
	var body = make(map[string]string)
	body["currentPage"] = "1"
	body["numPages"] = "10"
	body["pagerow"] = "16"
	body["totalPages"] = "5"
	body["filter_sel[cls1]"] = ""
	body["filter_sel[cls2]"] = ""
	body["order_sel"] = ""
	return body
}

func GetChineseFilmResponse(url string) {
	for index := 1; index < 17; index++ {
		body := fmt.Sprintf("currentPage=" + strconv.Itoa(index) + "&numPages=10&pagerow=16&totalPages=5&filter_sel%5Bcls1%5D=none&filter_sel%5Bcls2%5D=none&order_sel=none")
		response, err := downloader.PostHttpResponse(url, body, true)
		if err != nil {
			return
		}
		responseString := (*string)(unsafe.Pointer(&response))
		fmt.Println(responseString)
	}
}
func GetChineseFilmMethodTwo(url string) {
	for index := 1; index < 17; index++ {
		body := `{"currentPage":"` + strconv.Itoa(index) + `", "numPages":"10", "pagerow":"16", "totalPages":"5","filter_sel[cls1]":"none","filter_sel[cls2]":"none", "order_sel":"none"}`
		req := gorequest.New()
		resq, _, _ := req.Post(url).Set("X-Requested-With", "XMLHttpRequest").Send(body).End()
		resqByte, _ := ioutil.ReadAll(resq.Body)
		resqString := (*string)(unsafe.Pointer(&resqByte))
		fmt.Println(resqString)
	}
}
