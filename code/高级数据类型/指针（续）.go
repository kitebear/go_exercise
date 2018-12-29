package main

import "fmt"

type Pet interface {
	Name() string
	Age() uint8
}

type Dog struct {
	name string
	age  string
}

func (d Dog) Name() string {
	return d.name
}

func (d Dog) Age() uint8 {
	return d.age
}

func main() {
	myDog := Dog{"Little D", 3}
	_, ok1 := interface{}(&myDog).(Pet)
	_, ok2 := interface{}(myDog).(Pet)
	fmt.Printf("%v, %v\n", ok1, ok2)
}
