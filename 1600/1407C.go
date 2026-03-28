package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n)
	ans := make([]any, n+1)
	ans[0] = "!"
	i := 1
	for j := 2; j <= n; j++ {
		fmt.Println("?", i, j)
		fmt.Scan(&x)
		fmt.Println("?", j, i)
		fmt.Scan(&y)

		if x > y {
			ans[i] = x
			i = j
		} else {
			ans[j] = y
		}
	}
	ans[i] = n
	fmt.Println(ans...)
}
