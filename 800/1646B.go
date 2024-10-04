package main

import (
	"fmt"
	"sort"
)

func main() {
	var t, n int
next:
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		a := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])
		}
		sort.Ints(a)
		var pre, suf int64 = int64(a[0]), 0
		i, j := 1, n-1
		for i < j {
			pre += int64(a[i])
			suf += int64(a[j])
			if pre < suf {
				fmt.Println("YES")
				continue next
			}
			i++
			j--
		}
		fmt.Println("NO")
	}
}
