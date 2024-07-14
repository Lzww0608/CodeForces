package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a, b := n/7, n%7
	mn, mx := a*2, a*2+min(b, 2)
	if b == 6 {
		mn++
	}

	fmt.Println(mn, mx)
}
