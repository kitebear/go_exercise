package main

import (
	"fmt"
	"strconv"
	"sync/atomic"
)

// 员工ID生成器
type EmployeeIdGenerator func(company string, department string, sn uint32) string

// 默认公司名称
var company = "Gophers"

// 序列号
var sn uint32

// 生成员工ID
func generateId(generator EmployeeIdGenerator, department string) (string, bool) {
	// 这是一条 if 语句，我们会在下一章讲解它。
	// 若员工ID生成器不可用，则无法生成员工ID，应直接返回。
	if generator == nil {
		return "", false
	}
	// 使用代码包 sync/atomic 中提供的原子操作函数可以保证并发安全。
	newSn := atomic.AddUint32(&sn, 1)
	return generator(company, department, newSn), true
}

// 字符串类型和数值类型不可直接拼接，所以提供这样一个函数作为辅助。
func appendSn(firstPart string, sn uint32) string {
	return firstPart + strconv.FormatUint(uint64(sn), 10)
}

// 单返回值函数
func plus(a, b int) (res int) {
	return a + b
}

// 多返回值函数
func multi() (string, int) {
	return "name", 18
}

// 命名返回值
// 被命名的返回参数的值为该类型的默认零值
// 该例子中 name 默认初始化为空字符串，height 默认初始化为 0
func namedReturnValue() (name string, height int) {
	name = "xiaoming"
	height = 180
	return
}

// 参数可变函数
func sum(nums ...int) int {
	fmt.Println("len of nums is : ", len(nums))
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

func addInt(n int) func() int {
	i := 0
	return func() int {
		i += n
		return i
	}
}

func sayHello(name string) {
	fmt.Println("Hello ", name)
}

func logger(f func(string), name string) {
	fmt.Println("start calling method sayHello")
	f(name)
	fmt.Println("end calling method sayHellog")
}

func sendValue(name string) {
	name = "hemuketang"
}

func sendAddress(name *string) {
	*name = "hemuketang"
}

func main() {
	var generator EmployeeIdGenerator
	generator = func(company string, department string, sn uint32) string {
		return appendSn(company+"-"+department+"-", sn)
	}
	fmt.Println(generateId(generator, "RD"))

	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	// 匿名函数 立即执行函数
	func(name string) {
		fmt.Println(name)
	}("禾木课堂")

	// 闭包
	addTwo := addInt(2)
	fmt.Println(addTwo())
	fmt.Println(addTwo())
	fmt.Println(addTwo())

	// 函数作为参数  高阶函数
	logger(sayHello, "禾木课堂")

	// 传值和传引用
	str := "禾木课堂"
	fmt.Println("before calling sendValue, str : ", str)
	sendValue(str)
	fmt.Println("after calling sendValue, str : ", str)

	fmt.Println("before calling sendAddress, str : ", str)
	sendAddress(&str)
	fmt.Println("after calling sendAddress, str: ", str)
}
