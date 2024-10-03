package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		if n&1 == 1 {
			fmt.Print(1, " ")
		}
		for i := (n & 1) + 1; i <= n; i += 2 {
			fmt.Print(i+1, " ", i, " ")
		}
		fmt.Println()

	}
}
