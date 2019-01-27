package main

import "fmt"

type Pet interface {
	Name() string
	Age() uint8
}

type Dog struct {
	name string
	age  uint8
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Age() uint8 {
	return d.age
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	myDog := Dog{"Little D", 3}
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)

	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	fmt.Println(v)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
