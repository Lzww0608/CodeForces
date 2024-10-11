package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	m := map[int]string{
		0: "R",
		1: "O",
		2: "Y",
		3: "G",
		4: "B",
		5: "I",
		6: "V",
	}
	for i := 0; i < 7; i++ {
		n--
		fmt.Print(m[i])
	}
	idx := 0
	for n > 0 {
		n--
		fmt.Print(m[idx+3])
		idx = (idx + 1) % 4
	}
	fmt.Println()
}
