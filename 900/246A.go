package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 3 {
		fmt.Println(-1)
		return
	}
	for i := 0; i < n-1; i++ {
		fmt.Print(n-i, " ")
	}
	fmt.Println(1)
}
