package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// 函数式编程

func apply(lambda func(a, b int) int, a, b int) int {
	return lambda(a, b)
}

func main() {
	sub := func(a, b int) int {
		return a - b
	}
	println(apply(sub, 10, 8))

	println(apply(func(a, b int) int {
		return a + b
	}, 10, 8))

	println("-----------")
	a := f()
	for i := 0; i <= 100; i++ {
		fmt.Println(a(i))
	}

	println("----fib-------")
	// 将fib当成流输出
	printFileContent(fib())

}

// 闭包
func f() func(i int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

type intGen func() int

// 实现了这个接口，就能当做是文件一样的读
func (i intGen) Read(p []byte) (n int, err error) {
	next := i()
	if next > 100 {
		return 0, io.EOF
	}

	// 由于实现输出什么的比较复杂，所以这里借助一个已经实现好的结构体去做
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)

}

// 闭包fib
func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func printFileContent(reader io.Reader) {
	sc := bufio.NewScanner(reader)

	for sc.Scan() {
		println(sc.Text())
	}
}
