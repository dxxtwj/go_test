package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("https://jiuzan.plus/dx-api/task/receiveTask",
		url.Values{"id": {"337588156253884416"}})
	resp.Header.Add("scheme", "https")
	resp.Header.Add("accept", "application/json")
	resp.Header.Add("accept-encoding", "gzip,deflate,br")
	resp.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	resp.Header.Add("content-length", "21")
	resp.Header.Add("Content-Type", "application/json")
	resp.Header.Add("origin", "https://jiuzan.plus")
	resp.Header.Add("referer", "https://jiuzan.plus/")
	resp.Header.Add("sec-fetch-dest", "empty")
	resp.Header.Add("sec-fetch-mode", "cors")
	resp.Header.Add("user-agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Mobile Safari/537.36")
	resp.Header.Add("x-requested-with", "XMLHttpRequest")
	resp.Header.Add("x_token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2MTQ1MDM0NzAsInVzZXJuYW1lIjoiMjkyODY2NjQzNDM3MjYwODAwMTMyNDY1NTg0MTMifQ.3p2i3bbgDGsGgggUmJCJWM2Ca-6iI6_TTN5IykS0Hwc")
	resp.Header.Add("sec-fetch-mode", "cors")
	resp.Header.Add("sec-fetch-site", "same-site")
	//resp.Header.Add("aws-check", "false")
	if err != nil {
		// handle error
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		panic(err)
	}

	fmt.Println(string(body))
}
