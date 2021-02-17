package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Mobile Safari/537.36")
	client := http.Client{
		Transport:     nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 重定向的地址都是放在via里面的
			// 重定向内容放在request
			fmt.Println("redirect:", request, "via", via)
			return nil// 返回nil就是允许重定向
		},
		Jar:           nil,
		Timeout:       0,
	}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// httputil.DumpResponse 获取里面的内容
	s , err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", s)
}
