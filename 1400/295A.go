package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	queries := make([][3]int, m)
	for i := range queries {
		fmt.Fscan(in, &queries[i][0], &queries[i][1], &queries[i][2])
	}
	dif := make([]int, m+2)
	for i := 0; i < k; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		dif[l]++
		dif[r+1]--
	}

	diff := make([]int, n+2)
	sum := 0
	for i := 1; i <= m; i++ {
		sum += dif[i]
		x := sum * queries[i-1][2]
		diff[queries[i-1][0]] += x
		diff[queries[i-1][1]+1] -= x
	}

	sum = 0
	for i := 1; i <= n; i++ {
		sum += diff[i]
		fmt.Fprintf(out, "%d ", a[i-1]+sum)
	}
}
