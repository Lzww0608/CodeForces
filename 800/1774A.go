package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		var s string
		fmt.Scan(&s)
		ans := int(s[0] - '0')
		for _, c := range s[1:] {
			if c == '0' || ans < 0 {
				fmt.Print("+")
				ans += int(c - '0')
			} else {
				fmt.Print("-")
				ans -= 1
			}
		}
		fmt.Println()
	}
}
