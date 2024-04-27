package slices

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	id   int
	name string
}

func TestSlices(stdin *os.File) {
	scan := bufio.NewScanner(stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}

	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}
	fmt.Println("ss ", ss)

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	fmt.Println("ss ", ss)
}
