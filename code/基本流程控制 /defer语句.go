package main

import (
	"fmt"
)

func deferIt3() {
	f := func(i int) int {
		fmt.Printf("%d ", i)
		return i * 10
	}
	for i := 1; i < 5; i++ {
		defer fmt.Printf("%d ", f(i))
	}
}

func deferIt2() {
	for i := 1; i < 5; i++ {
		defer fmt.Print(i)
	}
}

func main() {
	// deferIt2()
	// deferIt4()

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", func(n int) int {
			fmt.Printf("%d ", n)
			return n
		}(fibonacci(i)))
	}
}

func deferIt4() {
	for i := 1; i < 5; i++ {
		defer func(n int) {
			fmt.Print(n)
		}(i)
	}
}

func fibonacci(num int) int {
	if num == 0 {
		return 0
	}
	if num < 2 {
		return 1
	}
	return fibonacci(num-1) + fibonacci(num-2)
}
