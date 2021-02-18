package scheduler

import "goTest/crawler/engine"

type QueuedScheduler struct {
	requestChan chan  engine.Request
	workerChan chan chan engine.Request// 有很多worker,每个worker有自己的chan
}

// 每个worker有自己的chan
func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

// 添加数据
func (s QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}
// 从外界告诉我们有一个worker可以负责接收request
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run()  {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var wokerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			// 如果同时有request、worker
			if len(requestQ) > 0 && len(wokerQ) > 0 {
				activeWorker = wokerQ[0]
				activeRequest = requestQ[0]
			}
			select {
				case r := <- s.requestChan://如果来的是request，那么我们就加进request队列中
					requestQ = append(requestQ, r)// 收到的数据缓存起来
				case w:= <- s.workerChan://如果来的是worker，那么我们就加进worker队列中
					wokerQ = append(wokerQ, w)// 收到的数据缓存起来
				case activeWorker <- activeRequest:// 同时有request和worker,那就把request发给这个worker
					wokerQ = wokerQ[1:]
					requestQ = requestQ[1:]
			}
		}
	}()
}

