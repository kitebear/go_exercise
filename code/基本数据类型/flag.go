package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

// 用命令行本目录下执行
// go run flag.go -name="Robert"
func init() {
	fmt.Println("--------init-------")
	// flag.StringVar(&name, "name", "everyone", "The greeting object.")
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")
	// var name = flag.String(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// 1
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question1")
	// 	flag.PrintDefaults()
	// }
	// 2
	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// flag.CommandLine.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }

	// flag.Parse()
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)
}
