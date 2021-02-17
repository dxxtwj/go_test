package main

import (
	"fmt"
	"sync"
)

// 返回值  chan <- 代表，这个方法，channel是送数据的
// 返回值   <-chan  代表这个方法是给数据外面的，就是只能获取
func createWorker(id int,wg * sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id,w)
	return w
}

type worker struct {
	in   chan int
	done func()
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)// 有20个任务
	for i, worker := range workers {
		worker.in <- 'a' + i // 分发数据给channels[i]
	}
	for i, worker := range workers {
		worker.in <- 'A' + i // 分发数据给channels[i]
	}
	wg.Wait()
}

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("www %d r %c\n", id, n)
		w.done()
	}
}
func main() {
	chanDemo()
}
