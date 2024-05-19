package main

import "fmt"

func main() {
	var t, n, m int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n, &m)
		fmt.Println(max(n, m))
	}
}
