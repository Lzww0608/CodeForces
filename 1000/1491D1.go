package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)
	fmt.Println((n - 1) >> 1)
	i, j := n-1, 0
	for k := range a {
		if k&1 == 0 {
			fmt.Print(a[i], " ")
			i--
		} else {
			fmt.Print(a[j], " ")
			j++
		}
	}
	fmt.Println()
}
