package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

type Student struct {
	Age           int
	Name, Address string
	// 结构体内嵌匿名成员  声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员
	Person
}

type Tree struct {
	value       int
	left, right *Tree
}

// type StructName struct {
// 	FieldName type
// }

func (person *Person) Move(address string) string {
	oldAddress := person.Address
	person.Address = address
	return oldAddress
}

func main() {
	p := Person{"Robert", "Male", 33, "Beijing"}
	oldAddress := p.Move("San Francisco")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
	stu := Student{Age: 20, Name: "new name", Address: "北京", Person: p}
	fmt.Println(stu)

	fmt.Println("stu.Age", stu.Age)
	fmt.Println("stu.Name", stu.Name)

	// 结构体中可以内嵌结构体
	// 但是需要注意的是：如果嵌入的结构体是本身，那么只能用指针。请看以下例子。
	tree := Tree{
		value: 1,
		left: &Tree{
			value: 1,
			left:  nil,
			right: nil,
		},
		right: &Tree{
			value: 1,
			left:  nil,
			right: nil,
		},
	}

	fmt.Printf(">>> %#v\n", tree)

	// 结构体是可以比较的
	// 前提是结构体中的字段类型是可以比较的
	tree1 := Tree{
		value: 2,
	}

	tree2 := Tree{
		value: 2,
	}

	fmt.Printf(">>> %#v\n", tree1 == tree2)
}
