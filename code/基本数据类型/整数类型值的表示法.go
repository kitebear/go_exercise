package main

import (
	"fmt"
)

func main() {
	var num1 int = -0x1000

	// 这里用到了字符串格式化函数。其中，%X用于以16进制显示整数类型值，%d用于以10进制显示整数类型值。
	fmt.Printf("16进制数 %X 表示的是 %d \n", num1, (num1))
}
