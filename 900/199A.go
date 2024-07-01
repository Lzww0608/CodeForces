package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 1 || n == 0 {
		fmt.Println(n, 0, 0)
		return
	}

	m := map[int]bool{
		0: true,
		1: true,
	}
	a, b := 1, 2
	for b <= n {
		m[b] = true
		if m[n-b] {
			fmt.Println(b, n-b, 0)
			return
		}
		a, b = b, a+b
	}

	fmt.Println("I'm too stupid to solve this problem")
}
