package main

import (
	hello "TheGoPath"
	"fmt"
	"os"
)

func main() {
	fmt.Println(hello.SayHello(os.Args[1:]))

}
