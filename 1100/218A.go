package main

import "fmt"

func main() {
	var n, k int

	fmt.Scan(&n, &k)
	a := make([]int, n*2+1)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= 2*n && k > 0; i += 2 {
		if a[i] > a[i-1]+1 && a[i] > a[i+1]+1 {
			a[i] -= 1
			k--
		}
	}

	for _, x := range a {
		fmt.Print(x, " ")
	}
	fmt.Println()
}
