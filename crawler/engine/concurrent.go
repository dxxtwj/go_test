package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)// 开始的东西，URL请求
	out := make(chan PaseResult)// 接收返回的东西
	e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)// fetch回来后，把返回来的url创建worker，并且把有价值的数据放out等等遍历
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		result := <-out// createWorker的out会到这里接收
		for _, item := range result.Items {
			log.Printf("got item: #%d : %v", itemCount, item)// 遍历请求回来的有价值的数据
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)// 当前请求回来的request再创建请求
		}
	}
}

func createWorker(in chan Request, out chan PaseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
