package engine

import (
	"goTest/crawler/fetcher"
	"log"
)

type SimpleEngine struct {

}

// 種子頁面，起始頁
func (e SimpleEngine) Run(seeds ...Request) {
	log.Println(1)
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//log.Println("requests" ,requests)//requests [{http://localhost:8080/mock/www.zhenai.com/zhenghun 0x64ab00}]
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("got item %v", item)
		}
	}
}

func worker(r Request) (PaseResult, error) {
	log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch:error"+"fetching url %s: %v", r.Url, err)
		return PaseResult{}, err
	}
	return r.ParseFunc(body), nil
}
