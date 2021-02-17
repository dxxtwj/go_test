package main

import (
	"fmt"
)

func main() {
	//c := make(chan int)// 开始的东西，URL请求
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	var a int
	a=<-pipe
	//mylib.ChanTest(c)
	//a :=  <- c
	fmt.Println(a)
	fmt.Println(pipe)
}
