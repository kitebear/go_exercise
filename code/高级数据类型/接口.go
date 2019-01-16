package main

import (
	"fmt"
	"math"
)

type Animal interface {
	Grow()
	Move(string) string
}

type Cat struct {
	Name    string
	Age     int8
	Address string
}

func (c *Cat) Grow() {
	c.Age++
}

func (c *Cat) Move(str string) string {
	return str
}

type tmp interface {
	area() float32
}

type geometry interface {
	// geometry 接口将会拥有 tmp 接口所定义的所有方法
	tmp
	perim() float32
}

type rect struct {
	len, wid float32
}

type circle struct {
	radius float32
}

func (r rect) area() float32 {
	return r.len * r.wid
}

func (r rect) perim() float32 {
	return 2 * (r.len + r.wid)
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func show(name string, param interface{}) {
	switch param.(type) {
	case geometry:
		// 类型断言
		fmt.Printf("area of %v is %v \n", name, param.(geometry).area())
		fmt.Printf("perim of %v is %v \n", name, param.(geometry).perim())
	default:
		fmt.Println("wrong type!")
	}
}

func main() {
	myCat := Cat{"Little C", 2, "In the house"}
	// 用来判断两者是否为实现关系，Cat实现了Animal
	animal, ok := interface{}(&myCat).(Animal)
	fmt.Printf("%v, %v\n", ok, animal)

	rec := rect{
		len: 1,
		wid: 2,
	}
	show("rect", rec)

	cir := circle{
		radius: 1,
	}
	show("circle", cir)

	show("test", "test param")
}
