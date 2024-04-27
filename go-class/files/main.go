package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}

		//data, err := io.ReadAll(file)
		//if err != nil {
		//	fmt.Fprint(os.Stderr, err)
		//	continue
		//}
		//fmt.Printf("The file has %d bytes \n", len(data))
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			s := scanner.Text()
			wc += len(strings.Fields(s))
			cc += len(s)
			lc++
		}
		fmt.Printf(" %7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}
}
