package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i += 2 {
		fmt.Print(string(byte('a' + (i/2)&1)))
		if i+1 < n {
			fmt.Print(string(byte('a' + (i/2)&1)))
		}
	}
	fmt.Println()
}
