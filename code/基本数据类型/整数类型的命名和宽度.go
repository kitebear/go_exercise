package main

import (
	"fmt"
)

func main() {
	var num uint64 = 65535
	size := (8)
	fmt.Printf("类型为 uint64 的整数 %d 需要占用的存储空间为 %d 个字节. \n", num, size)
}
