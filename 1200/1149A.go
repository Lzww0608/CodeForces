package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	a, b := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x == 1 {
			a++
		} else {
			b++
		}
	}

	if a == 0 || b == 0 {
		for b > 0 {
			fmt.Print(2, " ")
			b--
		}
		for a > 0 {
			fmt.Print(1, " ")
			a--
		}
		fmt.Println()
		return
	}
	fmt.Print(2, 1, " ")
	for b > 1 {
		fmt.Print(2, " ")
		b--
	}
	for a > 1 {
		fmt.Print(1, " ")
		a--
	}
	fmt.Println()
	return
}
