package main

import (
	"fmt"
	"math"
)


func Triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTringle(a, b))
}

func calcTringle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
