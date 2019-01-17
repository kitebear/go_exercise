package main

import (
	"fmt"
	"sync"
)

func main() {
	// ch1 := make(chan int, 1)
	// ch2 := make(chan int, 1)
	// ch3 := make(chan int, 3)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("1")
		// ch1 <- 1
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		// <-ch1
		// time.Sleep(100 * time.Millisecond)
		fmt.Println("2")
		wg.Done()
		// ch2 <- 2
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		// <-ch2
		// time.Sleep(200 * time.Millisecond)
		fmt.Println("3")
		wg.Done()
		// ch3 <- 3
	}()
	wg.Wait()
}
