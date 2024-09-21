package main

import "fmt"

func main() {
	var n int

	for i := 0; i < 1; i++ {
		fmt.Scan(&n)
		solve(n)
	}

	return
}

func solve(n int) {
	for i := 2; i <= n*2; i++ {
		if (i/2)*(i-i/2) >= n {
			fmt.Println(i)
			return
		}
	}
}
