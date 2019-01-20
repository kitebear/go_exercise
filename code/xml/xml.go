package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

type Stu struct {
	Name string `xml:"name"`
	// omitempty 忽略为空的标签
	Address string `xml:"addr,omitempty"`
	Secret  string `xml:"-"`
	Mother  string `xml:"parent>mother"`
	Father  string `xml:"parent>father"`
	Note    string `xml:"note,attr"`
	// Par     Parent `xml:"parent"`
}

// type Parent struct {
// 	Father string `xml:"father"`
// 	Mother string `xml:"mother"`
// }

func main() {
	stu := &Stu{
		Name:   "myName",
		Secret: "mySecret",
		Father: "myFather",
		Mother: "myMother",
		Note:   "MyNode",
		// Par: Parent{
		// 	Father: "myFather",
		// 	Mother: "myMonther",
		// },
	}

	//Block 1
	// data, err := xml.Marshal(stu)
	//Block 2
	data, err := xml.MarshalIndent(stu, "", "    ")
	errCheck(err)

	err = ioutil.WriteFile("stu.xml", data, 0644)
	errCheck(err)

	content, err := ioutil.ReadFile("stu.xml")
	errCheck(err)

	stuNew := &Stu{}
	xml.Unmarshal(content, stuNew)
	errCheck(err)

	fmt.Printf("content is : %#v\n", string(content))
	fmt.Printf("content is : %#v\n", stuNew)

	fmt.Println("stu.name : ", stuNew.Name)
}
