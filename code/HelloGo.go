// main.go
// 今天吧环境配了，并且debug环境也配好了
package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(runtime.GOOS)
	fmt.Println(1 << 3)
	var a []int
	append(a, 1)
	fmt.Println(a)
}

// func Pic(dx, dy int) [][]uint8 {
// 	var dt make([][]uint8, dx)
// 	for i:= 0; i<dx; i++ {
// 		dt[i] := make([]uint8, dy)
// 		for j:= 0; j<dy; j++ {
// 			dt[i][j] = (i+j)/2
// 		}
// 	}
// 	return dt
// }
