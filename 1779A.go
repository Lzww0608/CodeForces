package main

import "fmt"

func main() {
	var t, n int
next:
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		var s string
		fmt.Scan(&s)
		for i := 1; i < n; i++ {
			if s[i-1] == 'L' && s[i] == 'R' {
				fmt.Println(i)
				continue next
			} else if s[i-1] == 'R' && s[i] == 'L' {
				fmt.Println(0)
				continue next
			}

		}
		fmt.Println(-1)
	}
}
