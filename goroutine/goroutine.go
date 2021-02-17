package main

import (
	"fmt"
	"time"
)

func main() {
	for i:= 0; i < 1000; i++ {
		go func(i int) {// 并发去开了这个函数
			for {
				fmt.Printf("hhh %d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}