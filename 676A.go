package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	mn, mx := -1, -1
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == 1 {
			mn = i
		} else if x == n {
			mx = i
		}
	}
	fmt.Println(max(mx, mn, n-1-mx, n-1-mn))
}
