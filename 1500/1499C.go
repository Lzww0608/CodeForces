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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	mn := [2]int{a[0], a[1]}
	ans := math.MaxInt
	sum := [2]int{0, 0}
	cnt := [2]int{}
	for i, x := range a {
		p := i % 2
		cnt[p]++
		sum[p] += x
		mn[p] = min(mn[p], x)
		if i > 0 {
			cur := sum[0] + sum[1] + mn[0]*(n-cnt[0]) + mn[1]*(n-cnt[1])
			ans = min(ans, cur)
		}
	}

	fmt.Fprintln(out, ans)
}
