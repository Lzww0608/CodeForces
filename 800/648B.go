package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n*2)
	for i := range a {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	for i, j := 0, 2*n-1; i < j; i, j = i+1, j-1 {
		fmt.Println(a[i], a[j])
	}

}
