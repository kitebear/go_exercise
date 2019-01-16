package main

import "fmt"

func main() {
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	// go语言中 = 是赋 而  := 是声明变量并赋值。
	slice3 := numbers3[2:len(numbers3)]
	length := (3)
	capacity := (3)
	fmt.Printf("%v, %v\n", (length == len(slice3)), (capacity == cap(slice3)))

	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println("len: ", len(slice))
	// fmt.Println("cap: ", cap(slice))

	// //改变切片长度
	// slice = append(slice, 6)
	// fmt.Println("after append operation: ")
	// fmt.Println("len: ", len(slice))
	// fmt.Println("cap: ", cap(slice)) //注意，底层数组容量不够时，会重新分配数组空间，通常为两倍

	// 多个切片共享一个底层数组的情况
	slice := []int{1, 2, 3, 4, 5}
	// newSlice := slice[0:3]
	// 使用 copy 方法可以避免共享同一个底层数组
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)

	fmt.Println("before modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println()

	// newSlice[0] = 6
	// fmt.Println("after modifying underlying array:")
	// fmt.Println("slice: ", slice)
	// fmt.Println("newSlice: ", newSlice)
}
