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

	cnt, ans := 0, n
	for i := 1; i <= a[n-1]; i <<= 1 {
		cur := 0
		for _, x := range a {
			if x%i == 0 {
				cur++
			}
		}
		if cur > 0 {
			ans, cnt = i, cur
		}
	}
	fmt.Println(ans, cnt)
}
