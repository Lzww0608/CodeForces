package main

import (
	"fmt"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	x--
	ans := a[x]
	for i, j := x-1, x+1; i >= 0 || j < n; i, j = i-1, j+1 {
		tmp := 0
		if i >= 0 && a[i] == 1 {
			tmp++
		}
		if j < n && a[j] == 1 {
			tmp++
		}

		if tmp == 2 {
			ans += 2
		} else if tmp == 1 && (i < 0 || j >= n) {
			ans += 1
		}

	}

	fmt.Println(ans)
}
