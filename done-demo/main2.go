package main

import (
	"fmt"
	"time"
)

// / work pool
func worker(id int, jobs <-chan int, results chan<- int, done chan bool) {
	for job := range jobs {
		fmt.Printf("workder: %d start job: %d\n", id, job)
		results <- job * 2
		time.Sleep(time.Second)
		fmt.Printf("workder: %d stop job: %d\n", id, job)
	}
	done <- true
}
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	done := make(chan bool)
	//开启三个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results, done)
	}

	//发送五个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for {
		select {
		case n := <-results:
			fmt.Println(n)
		case <-done:
			close(results)
			fmt.Println("ok")
			return
		}
	}

	//for tmp := range results {
	//	fmt.Println(tmp)
	//}
}
