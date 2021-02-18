package parser

import (
	"goTest/crawler/engine"
	"log"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
// 城市列表解析器
func ParsetCityList(contents []byte) engine.PaseResult {
	log.Println(2)
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.PaseResult{}
	for _, m := range matches {// 打開城市中每一頁的用戶數據
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}
