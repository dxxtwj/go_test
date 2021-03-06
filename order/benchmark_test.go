package benchmark

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func QSort(data []int) {
	myqsort(data, 0, len(data)-1)
}
var ints []int

func myqsort(data []int, s, e int) {
	if s >= e {
		return
	}

	t := data[s]
	i, j := s, e

	for i < j {
		for ; i < j && data[j] >= t; j-- { }
		for ; i < j && data[i] < t; i++ { }
		if i < j { break }

		data[i], data[j] = data[j], data[i]
		i++
		j--
	}

	data[i] = t
	myqsort(data, s, i-1)
	myqsort(data, i+1, e)
}


// 长度为 1w 的数据使用系统自带排序
func BenchmarkSort10k(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()   // 只考虑下面代码的运行事件，所以重置计时器
	for i := 0; i < t.N; i++ {
		sort.Ints(slice)
	}
}

// 长度为 100 的数据使用系统自带排序
func BenchmarkSort100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		sort.Ints(slice)
	}
}

// 长度为 1w 的数据使用上述代码排序
func BenchmarkQsort10k(t *testing.B) {
	slice := ints[0:10000]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		QSort(slice)
	}
}

// 长度为 100 的数据使用上述代码排序
func BenchmarkQsort100(t *testing.B) {
	slice := ints[0:100]
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		QSort(slice)
	}
}

// 数据初始化，为了保证每次数据都是一致的。
func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	ints = make([]int, 10000)

	for i := 0; i < 10000; i++ {
		ints[i] = rand.Int()
	}
	m.Run()
}