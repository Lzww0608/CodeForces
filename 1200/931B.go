package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	a, b = max(a, b)-1, min(a, b)-1

	c := a ^ b
	ans := 0
	for c > 0 {
		ans++
		c >>= 1
	}

	if 1<<ans == n {
		fmt.Println("Final!")
	} else {
		fmt.Println(ans)
	}

}
