package main

import "fmt"

func main() {
	var t, n, m int

	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n, &m)
		if n == 1 && m == 1 {
			fmt.Println(n + m - 2)
			continue
		}
		fmt.Println((n + m) - 2 + min(m, n))

	}
}
