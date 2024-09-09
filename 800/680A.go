package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 5)
	sum := 0
	for i := range a {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	sort.Ints(a)
	ans := 0
	for i := 0; i < 4; i++ {
		if a[i] == a[i+1] {
			ans = max(ans, a[i]*2)
			if i < 3 && a[i] == a[i+2] {
				ans = max(ans, a[i]*3)
			}
		}
	}

	fmt.Println(sum - ans)
}
