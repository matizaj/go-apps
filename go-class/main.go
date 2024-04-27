package main

import (
	"go-class/slices"
	"log"
	"os"
)

func main() {
	f, err := os.Open("text.txt")
	if err != nil {
		log.Panicln("cant open file", err)
	}
	slices.TestSlices(f)
}
