package main

import (
	"fmt"
	"sort"
)

func main() {
	var t, n, l, r, k int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n, &l, &r, &k)
		a := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])
		}
		sort.Ints(a)
		ans := 0
		for _, x := range a {
			if k-x >= 0 && x >= l && x <= r {
				ans++
				k -= x
			} else if x > r {
				break
			}
		}
		fmt.Println(ans)
	}
}
