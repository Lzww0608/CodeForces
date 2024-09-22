package main

import (
	"fmt"
	"strings"
)

func main() {
	var k, d int
	fmt.Scan(&k, &d)
	if d == 0 {
		if k == 1 {
			fmt.Println(0)
		} else {
			fmt.Println("No solution")
		}
		return
	} else if d == 9 {
		fmt.Println(strings.Repeat("9", k))
		return
	} else {
		fmt.Print(d)
		if k > 1 {
			fmt.Println(strings.Repeat("9", k-1))
		}
	}
}
