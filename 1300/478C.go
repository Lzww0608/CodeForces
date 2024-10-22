package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a, b := n/m, n%m
	mn := a*(a-1)/2*(m-b) + a*(a+1)/2*b
	mx := (n - m + 1) * (n - m) / 2
	fmt.Println(mn, mx)
}
