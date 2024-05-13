package main

import "fmt"

func main() {
	var t, n, k int
	for fmt.Scan(&t); t > 0; t-- {
		fmt.Scan(&n, &k)
		if k != 1 && k > (n+1)/2 {
			fmt.Println(-1)
			continue
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j && i&1 == 0 && k > 0 {
					fmt.Print("R")
					k--
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}
