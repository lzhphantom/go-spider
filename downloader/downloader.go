package downloader

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

func GetHttpResponse(url string, ok bool) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("User_Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, err
	}

	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	} else {
		return ioutil.ReadAll(response.Body)
	}
}
