package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		for i := 1; i <= n; i++ {
			fmt.Print(i*2, " ")
		}
		fmt.Println()
	}
}
