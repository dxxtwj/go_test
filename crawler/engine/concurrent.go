package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request // 我有一个worker，请问给我哪一个chan?
	Run()
}
type ReadyNotifier interface {
	WorkerReady(w chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan PaseResult) // 接收返回的东西
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		// 问e.Scheduler.WorkerChan()要一个chan
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler) // fetch回来后，把返回来的url创建worker，并且把有价值的数据放out等等遍历
	}
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("Duplicate request:"+"%s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out // createWorker的out会到这里接收
		for _, item := range result.Items {
			go func() {e.ItemChan <- item}()
		}
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("Duplicate request:"+"%s", request.Url)
				continue
			}
			e.Scheduler.Submit(request) // 当前请求回来的request再创建请求
		}
	}
}

func createWorker(in chan Request, out chan PaseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request) // 做事情
			if err != nil {
				continue
			}
			out <- result // 做完事情要输出了
		}
	}()
}

var visitedUrls = make(map[string]bool)

// 判断是否有重复
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = true
	return false
}
