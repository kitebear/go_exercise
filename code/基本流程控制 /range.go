package main

import "fmt"

func main() {
	var str = "hello, 世界"

	for _, char := range str {
		fmt.Printf("%T", char)
	}
}
