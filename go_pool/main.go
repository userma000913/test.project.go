package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

func main() {
	p, _ := ants.NewPool(100000, ants.WithPreAlloc(true))
	defer p.Release()
	for {
		p.Submit(func() {
			fmt.Println("打印")
			time.Sleep(time.Second)
		})
	}
}
