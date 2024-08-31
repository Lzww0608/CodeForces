package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	a := make([]int, m+1)
	for i := range a {
		fmt.Scan(&a[i])
	}

	ans := 0
	for i := 0; i < m; i++ {
		x := a[i] ^ a[m]
		if bits.OnesCount(uint(x)) <= k {
			ans++
		}
	}

	fmt.Println(ans)
}
