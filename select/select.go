package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("www %d r %d\n", id, n)
	}
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): // 800毫秒内没有生成数据访问这里
			fmt.Println("timeout")
		case <-tick:// 每秒钟看一下队列的长度
			fmt.Println("queue ;len = ", len(values))
		case <-tm: // 达到多少秒后结束
			fmt.Println("bye")
			return
		}
	}
}
