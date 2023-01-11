package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

// 定时任务demo
func main() {
	go f()
	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	//for {
	//	s := <-sigChan
	//	log.Printf("get a signal %s\n", s.String())
	//	switch s {
	//	case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
	//		log.Println("metis.statistics.server server exit now...")
	//		return
	//	case syscall.SIGHUP:
	//	default:
	//	}
	//}

	// 阻塞主程序，防止停止
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			// 续期
			t1.Reset(time.Second * 10)
		}
	}
}

/*
*
http://cron.ciding.cc/ cron表达式生成器
*/
func f() {
	c := cron.New()
	c.AddFunc("0/1 * 0/1 * * ?", func() {
		fmt.Println("定时任务执行了")
	})
	c.Start()
}
