package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int, 10)
	worker := 9
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i < worker; i++ {
		go work(i, jobs, &wg)
	}
	AddJob(jobs, 10)
	wg.Wait()
}

func AddJob(jobs chan int, len int) {
	for i := 0; i < len; i++ {
		jobs <- i
	}
}

func work(i int, jobs chan int, wg *sync.WaitGroup) {
	for job := range jobs {
		wg.Done()
		fmt.Printf("worker is %d, job is %d \n", i, job)
		time.Sleep(time.Second)
	}

}
