package main

import "fmt"

func main() {
	var t, n, m int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n, &m)
		a := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])
		}
		s := make([]byte, m)
		for i := 0; i < n; i++ {
			x, y := min(a[i]-1, m+1-a[i]-1), max(a[i]-1, m+1-a[i]-1)
			if s[x] != 'A' {
				s[x] = 'A'
			} else {
				s[y] = 'A'
			}
		}
		for i := range s {
			if s[i] == 'A' {
				fmt.Print("A")
			} else {
				fmt.Print("B")
			}
		}
		fmt.Println()
	}
}
