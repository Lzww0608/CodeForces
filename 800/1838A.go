package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		var mx, mn int = -1e9, 1e9
		var x int
		for i := 0; i < n; i++ {
			fmt.Scan(&x)
			mx = max(mx, x)
			mn = min(mn, x)
		}
		if mn < 0 {
			fmt.Println(mn)
		} else {
			fmt.Println(mx)
		}

	}
}
