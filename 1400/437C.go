package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	v := make([]int, n)
	for i := range v {
		fmt.Scan(&v[i])
	}
	ans := 0
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		ans += min(v[a-1], v[b-1])
	}
	fmt.Println(ans)
}
