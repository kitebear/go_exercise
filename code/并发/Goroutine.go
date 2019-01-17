package main

import (
	"fmt"
	"log"
	"time"
)

func doSomething(id int) {
	log.Printf("before do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d) \n", id)
}

/**
并发指的是多个任务被（一个）cpu 轮流切换执行，在 Go 语言里面主要用 goroutine （协程）来实现并发，类似于其他语言中的线程（绿色线程）。
多个 goroutine 的执行是随机。
对于 IO 密集型任务特别有效，比如文件，网络读写。
*/
func main() {
	go doSomething(1)
	go doSomething(2)
	go doSomething(3)

	time.Sleep(4 * time.Second)

	for i := 0; i < 3; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
