package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	x := make([]int, n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x[i], &h[i])
	}

	ans := 1
	pre := math.MinInt32
	for i := 0; i < n-1; i++ {
		if x[i]-h[i] > pre {
			ans++
			pre = x[i]
		} else if x[i]+h[i] < x[i+1] {
			ans++
			pre = x[i] + h[i]
		}
		pre = max(pre, x[i])
	}

	fmt.Fprintln(out, ans)
}
