package main

import (
	hello "TheGoPath"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(hello.SayHello(os.Args[1]))
	} else {
		fmt.Println("Hello, goblin!")
	}
}
