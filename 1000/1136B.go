package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	ans := n*2 + 1
	ans += min(k-1, n-k)*2 + max(k-1, n-k)
	fmt.Println(ans)
}
