package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		a := make([]int, n)
		for i := range a {
			a[i] = i + 1
		}
		fmt.Println(n)
		for i := 0; i < n; i++ {
			for _, x := range a {
				fmt.Print(x, " ")
			}
			fmt.Println()
			if i < n-1 {
				a[n-1], a[n-i-2] = a[n-i-2], a[n-1]
			}

		}

	}
}
