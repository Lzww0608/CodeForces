package main

import "fmt"

func main() {
	var t, n int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n)
		if n < 100 {
			fmt.Println(n % 10)
		} else {
			mn := 9
			for n > 0 {
				mn = min(mn, n%10)
				n /= 10
			}
			fmt.Println(mn)
		}
	}
}
