package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n<<1)
	for i := range a {
		fmt.Scan(&a[i])
	}

	sort.Ints(a)
	if a[0] == a[n*2-1] {
		fmt.Println(-1)
	} else {
		for _, x := range a {
			fmt.Print(x, " ")
		}
		fmt.Println()
	}

}
