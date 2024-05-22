package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		a := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])
		}
		for i := n - 1; i >= 0; i-- {
			fmt.Print(a[i], " ")
		}
		fmt.Println()
	}
}
