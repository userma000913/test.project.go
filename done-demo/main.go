package main

import (
	"fmt"
)

// 使用done通道来控制   和 waitGroup一样
func receive(c chan int, done chan bool) {
	for item := range c {
		fmt.Println(item)
	}
	done <- true
}

func send(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c)

}
func main() {
	c := make(chan int, 3)
	done := make(chan bool)
	go send(c)
	go receive(c, done)
	<-done
}
