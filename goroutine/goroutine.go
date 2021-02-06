package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i:= 0; i < 10; i++ {
		go func(ii int) {// 并发去开了这个函数
			for {
				a[ii]++
				runtime.Gosched()// 交出控制权
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}