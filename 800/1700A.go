package main

import "fmt"

func main() {
	var t, n, m int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n, &m)
		ans := m * (m + 1) / 2
		a, b := m*2, m*n
		ans += (a + b) * (n - 1) / 2
		fmt.Println(ans)
	}
}
