package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n/2 + 1)
	for i := 1; i <= n; i += 2 {
		for j := 0; j < min(2, n-i+1); j++ {
			fmt.Println((i-1)/2+1, (i-1)/2+1+j)
		}

	}
}
