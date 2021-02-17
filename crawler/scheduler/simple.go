package scheduler

import "goTest/crawler/engine"

type SimpleScheduler struct {
	worekerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.worekerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.worekerChan <- r
	}() // 当前的scheduler
}
