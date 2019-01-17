package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
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

func panicFunc() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("recover panic")
		}
	}()

	panic(errors.New("this is test for panic"))
}

func main() {
	fmt.Println("before panic")

	panicFunc()

	fmt.Println("after panic")
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

// 使用 defer 关闭 http 请求响应体的 Body
func closeBody(url string) error {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	// ... do more stuff ...

	return err
}

// 使用 defer 关闭文件句柄
func closeFile(filename string) error {
	f, err := os.Open(filename)
	defer f.Close()
	// ... do more stuff ...

	return err
}

func mutex () float32 {
	return 1.32
}

type Customer struct {
	mutex() float32
}

// 使用 defer 解锁
func BillCustomer(c *Customer) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// ... do more stuff ...

	return
}

func panicFunc() {
	panic(errors.New("this is test for panic"))
}

// 最终返回的是 4。因为return 2 执行后，变量 i 赋值为 2， 但是随后执行了 defer 函数，i 被赋值为4，所以最终返回结果为4。
func testDefer() (i int) {
	defer func() {
		fmt.Println(i)
		i = 4
	}()

	return 2
}
