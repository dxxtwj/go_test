package mylib

func Add(a, b int) int {
	return a + b
}

func ChanTest(c chan int) {
	c <- 1
}