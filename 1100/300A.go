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
	fmt.Print(1, " ")
	fmt.Println(a[0])
	if a[n-1] > 0 {
		fmt.Print(1, " ")
		fmt.Println(a[n-1])
		fmt.Print(n-2, " ")
		for i := 1; i < n-1; i++ {
			fmt.Print(a[i], " ")
		}
	} else {
		fmt.Print(2, " ")
		fmt.Println(a[1], a[2])
		fmt.Print(n-3, " ")
		for i := 3; i < n; i++ {
			fmt.Print(a[i], " ")
		}
	}

	fmt.Println()
}
