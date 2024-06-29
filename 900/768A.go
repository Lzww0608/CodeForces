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
	cnt := n
	for i := 0; i < n; i++ {
		if a[i] == a[0] {
			cnt--
		} else {
			break
		}
	}

	for i := n - 1; i >= 0 && cnt > 0; i-- {
		if a[i] == a[n-1] {
			cnt--
		} else {
			break
		}
	}

	fmt.Println(cnt)
}
