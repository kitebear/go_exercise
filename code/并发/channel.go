package main

import (
	"fmt"
	"time"
)

// goroutine 是 Go 中实现并发的重要机制，channel 是 goroutine 之间进行通信的重要桥梁。
// 使用内建函数 make 可以创建 channel，举例如下：
// ch := make(chan int)  // 注意： channel 必须定义其传递的数据类型
// var ch chan int
/**

ch <- x // channel 接收数据 x
x <- ch // channel 发送数据并赋值给 x
<- ch // channel 发送数据，忽略接受者


上文提到，可以通过 make(chan int) 创建channel，此类 channel 称之为非缓冲通道。事实上 channel 可以定义缓冲大小，如下：
chInt := make(chan int)       // unbuffered channel  非缓冲通道
chBool := make(chan bool, 0)  // unbuffered channel  非缓冲通道
chStr := make(chan string, 2) // bufferd channel     缓冲通道
*/

func receive(receiver chan<- string, msg string) {
	receiver <- msg
}

func send(sender <-chan string, receiver chan<- string) {
	msg := <-sender
	receiver <- msg
}

func strWorker(ch chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("do something with strWorker...")
	ch <- "str"
}

func intWorker(ch chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("do something with intWorker...")
	ch <- 1
}

func worker(done chan bool) {
	fmt.Println("start working...")
	done <- true
	fmt.Println("end working...")
}

func main() {
	// Block 1
	// ch := make(chan string)
	// go func() {
	// 	fmt.Println("111")
	// 	ch <- "ping"
	// 	time.Sleep(2 * time.Second)
	// }()
	// fmt.Println("222")
	// fmt.Println(<-ch)

	// Block 2
	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2

	// ch1 := make(chan string, 1)
	// ch2 := make(chan string, 1)

	// receive(ch1, "pass message")
	// send(ch1, ch2)

	// fmt.Println(<-ch2)

	// Block 3
	// ch := make(chan int, 10)
	// for i := 0; i < 10; i++ {
	// 	ch <- i
	// }
	// close(ch)

	// res := 0
	// for v := range ch {
	// 	res += v
	// }

	// for v := range ch {
	// 	res += v
	// }

	// fmt.Println(res)

	// Block 4
	// chStr := make(chan string)
	// chInt := make(chan int)

	// go strWorker(chStr)
	// go intWorker(chInt)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case <-chStr:
	// 		fmt.Println("get value from strWorker")

	// 	case <-chInt:
	// 		fmt.Println("get value from intWorker")

	// 	}
	// }

	done := make(chan bool, 1)

	go worker(done)
	fmt.Println("<- done")
	<-done
}
