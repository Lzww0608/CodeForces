package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	fmt.Println(solve(n, k, (1<<n)-1))
}

func solve(n, k, t int) int {
	if k == (t+1)/2 {
		return n
	} else if k > (t+1)/2 {
		k -= (t + 1) / 2
	}

	return solve(n-1, k, t/2)
}
