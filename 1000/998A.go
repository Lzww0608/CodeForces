package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	if n == 1 || (n == 2 && a[0] == a[1]) {
		fmt.Println(-1)
	} else {
		fmt.Println(1)
		ans, mn := 0, a[0]
		for i, x := range a {
			if x < mn {
				mn, ans = x, i
			}
		}
		fmt.Println(ans + 1)
	}
}
