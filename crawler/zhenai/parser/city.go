package parser

import (
	"goTest/crawler/engine"
	"regexp"
)

// <a href="http://localhost:8080/mock/album.zhenai.com/u/7549966999891941491">心事痕迹万能萌妹</a>
const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

// 城市解析器
func ParseCity(contents []byte) engine.PaseResult {
	//log.Println(3)
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.PaseResult{}
	for _, m := range matches { // 把每一頁的用戶數據全部追加返回，目的是在engine.go中把所有都遍歷出來
		name := string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(c []byte) engine.PaseResult {
				return ParseProfile(c, name)
			},
		})

	}
	return result
}
