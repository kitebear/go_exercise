package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var (
		total1 int32 = 0
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				// atomic 自带原子操作
				atomic.AddInt32(&total1, 1)
				// fmt.Println("now total: ", total1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	if atomic.CompareAndSwapInt32(&total1, 0, 100) {
		fmt.Println("compare and swap successfully")
	}

	// fmt.Println("old valueis: ", atomic.SwapInt32(&total1, 200))

	fmt.Println("before atomic.Store, the value is :", total1)

	atomic.StoreInt32(&total1, 100)

	fmt.Println("The total number is", total1)
}
