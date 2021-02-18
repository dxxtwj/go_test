package scheduler

import "goTest/crawler/engine"

type SimpleScheduler struct {
	worekerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.worekerChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.worekerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.worekerChan <- r
	}() // 当前的scheduler
}
