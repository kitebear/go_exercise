package main

import "fmt"

func main() {
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	// go语言中 = 是赋 而  := 是声明变量并赋值。
	slice3 := numbers3[2:len(numbers3)]
	length := (3)
	capacity := (3)
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))
}
