package main

import (
	"goTest/crawler/engine"
	"goTest/crawler/scheduler"
	"goTest/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount:100,
	}
	//e.Run(engine.Request{
	//	Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParsetCityList,
	//})
	e.Run(engine.Request{
		Url:       "http://localhost:8080/mock/www.zhenai.com/zhenghun/ningde",
		ParseFunc: parser.ParseCity,
	})
}
