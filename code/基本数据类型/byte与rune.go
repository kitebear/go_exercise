package main

import (
	"fmt"
)

func main() {
	var char1 rune = '赞'

	// 这里用到了字符串格式化函数。其中，%c用于显示rune类型值代表的字符。
	fmt.Printf("字符 '%c' 的Unicode代码点是 %s。\n", char1, ("U+8D5E"))
}
