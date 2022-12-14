package main

import "fmt"

// 该例子用来演示channel是引用类型最好的例子！
func main() {

	f()

}

func f() {
	c := make(chan int)
	go f1(c)
	fmt.Println(<-c)
}
func f1(c chan int) {
	var c1 chan int

	c2 := make(chan int, 1)
	c2 <- 2

	// 指针拷贝 将c1的指针指向了c
	c1 = c
	// c1的改动会引起c的改动
	c1 <- <-c2
}
