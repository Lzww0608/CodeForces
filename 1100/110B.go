package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print(string('a' + i%4))
	}
	fmt.Println()
}
