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
	for i := 0; i+2 < n; i++ {
		if a[i]+a[i+1] > a[i+2] {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
