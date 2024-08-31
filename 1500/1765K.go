package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	sum, mn := 0, 1_000_000_000
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&x)
			sum += x
			if i+j == n-1 {
				mn = min(mn, x)
			}
		}
	}
	fmt.Println(sum - mn)
}
