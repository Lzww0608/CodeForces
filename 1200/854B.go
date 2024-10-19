package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	if k == 0 || k == n {
		fmt.Println(0, 0)
		return
	}
	fmt.Println(1, min(k*2, n-k))
}
