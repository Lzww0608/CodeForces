package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	n := a + b + 1
	for i := 1; i <= a; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print(n, " ")
	for i := 0; i < b; i++ {
		fmt.Print(n-1, " ")
		n--
	}
	fmt.Println()
}
