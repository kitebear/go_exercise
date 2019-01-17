package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// var (
	// wg      sync.WaitGroup
	// numbers []int
	// sync.RWMutex 的读的并发能力大概是 sync.Mutex 的十倍，从而大大提高了其并发能力。
	// rw      sync.RWMutex
	// mux sync.Mutex
	// )

	// Block 1
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		numbers = append(numbers, i)
	// 		wg.Done()
	// 	}(i)
	// }

	// Block 2
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		mux.Lock()
	// 		numbers = append(numbers, i)
	// 		mux.Unlock()
	// 		wg.Done()
	// 	}(i)
	// }

	// wg.Wait()

	// fmt.Println("The numbers is", numbers)

	/**
	我们可以通过共享内存的方式实现多个 goroutine 中的通信。
	多个 goroutine 对于共享的内存进行写操作的时候，可以使用 Lock 来避免数据不一致的情况。
	对于可以分离为读写操作的共享数据可以考虑使用 sync.RWMutex 来提高其读的并发能力。
	*/

	var (
		total = 0
		mux   sync.Mutex
	)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				mux.Lock()
				total += 1
				fmt.Println("total: ", total)
				time.Sleep(time.Millisecond)
				mux.Unlock()
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("finally total is: ", total)
}
