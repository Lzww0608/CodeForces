package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		if n&1 == 1 {
			fmt.Println(-1)
		} else {
			fmt.Println(0, n>>1, n>>1)
		}
	}
}
