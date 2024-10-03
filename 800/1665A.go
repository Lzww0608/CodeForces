package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		fmt.Println(n-3, 1, 1, 1)

	}
}
