package url

import (
	"fmt"
	"time"
)

var (
	BoxOfficeURL     = "http://www.piaofang.biz/"
	ChineseBoxOffice = "http://www.cbooo.cn/"
	CBooo_Movie      = "http://www.cbooo.cn/movies"
	Cbooo_RealUrl    = "http://www.cbooo.cn/Mdata/getMdata_movie?area=%s&type=0&year=0&initial=全部&pIndex=%d"
	MaoYanByYear     = "https://piaofang.maoyan.com/rankings/year?year=%d&limit=100&tab=%d"
)

func GetMaoYanUrl(year int) string {
	nowYear := time.Now().Year()
	tab := nowYear - year + 1
	return fmt.Sprintf(MaoYanByYear, year, tab)
}
