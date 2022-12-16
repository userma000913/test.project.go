package test

import (
	"fmt"
	"math"
	"testing"
)

// 运行键的第三个是输出函数的覆盖率
func TestAdd(t *testing.T) {
	test := []struct {
		a, b, c int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{11, 12, 23},
		{0, 0, 0},
	}

	for _, tt := range test {
		if sum := Add(tt.a, tt.b); sum != tt.c {
			t.Errorf("err,sum should is %d\n", sum)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	aa := math.MaxInt
	bb := 1
	c := math.MinInt

	// 系统默认循环次数
	for i := 0; i < b.N; i++ {
		if sum := Add(aa, bb); sum != c {
			b.Errorf("err,sum should is %d\n", sum)
		}
	}

	/**
	goos: darwin
	goarch: arm64
	pkg: test.project/test-demo/test
	BenchmarkAdd
	BenchmarkAdd-10    	1000000000	         0.3189 ns/op
	PASS
	*/

	//   执行1000000000次 一次执行耗费实现 0.3189纳秒

}

// go test bench . -cpuprofile cpu.out   // 测试文件目录下执行
// go tool pprof cpu.out
// web
// quit 退出

func ExampleAdd() {
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1, 2))

	// Output
	// 2
}
