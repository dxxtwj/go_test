package main

import (
	"fmt"
	"time"
)

// 返回值  chan <- 代表，这个方法，channel是送数据的
// 返回值   <-chan  代表这个方法是给数据外面的，就是只能获取
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i // 分发数据给channels[i]
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i // 分发数据给channels[i]
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) // 第二个参数是缓冲区大小
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	time.Sleep(time.Millisecond)
}

func worker(id int, c chan int) {
	// 判断管道是否有close了
	// 第一种方法，通过range去判断是否有close
	for n := range c{
		// 第一种方法判断close 如果close了就停止
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("www %d r %d\n", id, n)
	}
}

func channelClose() {
	c := make(chan int) // 第二个参数是缓冲区大小
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)// 告诉接受方，我发完数据了,执行到这里，后面的全部都是空数据的接受，通过n, ok := <-c可以做判断是否close

	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("channel as firest-class citizen")
	chanDemo()
	fmt.Println("buffered channel")
	bufferedChannel()
	fmt.Println("channel close and range")
	channelClose()
}

