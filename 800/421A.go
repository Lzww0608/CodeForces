package main

import "fmt"

func main() {
	var n, a, b, x int
	fmt.Scan(&n, &a, &b)
	m := map[int]struct{}{}
	for i := 0; i < a; i++ {
		fmt.Scan(&x)
		m[x] = struct{}{}
	}
	for i := 0; i < b; i++ {
		fmt.Scan(&x)
	}

	for i := 1; i <= n; i++ {
		if _, ok := m[i]; ok {
			fmt.Print(1, " ")
		} else {
			fmt.Print(2, " ")
		}
	}
	fmt.Println()
}
