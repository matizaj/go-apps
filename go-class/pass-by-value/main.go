git addpackage main

import "fmt"

func main() {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	fmt.Println("m", m)
	v := do(m)
	fmt.Println("m", m)
	fmt.Println("V", v)
}

func do(m1 map[int]int) int {
	m1[0] = 0
	fmt.Println("m1", m1)
	return m1[3]
}
