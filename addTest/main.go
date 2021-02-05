package main

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	tests := []struct {a, b, c int} {
		{3,4,5},
		{5,12,13},
		{8,15,17},
		{12,35,37},
		{30000,40000,50000},
	}
	for _, tt := range tests {
		if actual := calcTringle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTringle(%d, %d);" + "got %d; expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

func main() {
	for i:= 0 ; i > 100 ; i++ {
		fmt.Println(1)
	}
}
