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
	mn, mx := math.MaxInt32, math.MinInt32
	for i, x := range a {
		if x != -1 {
			continue
		}
		if i > 0 && a[i-1] != -1 {
			mn = min(mn, a[i-1])
			mx = max(mx, a[i-1])
		}
		if i+1 < n && a[i+1] != -1 {
			mn = min(mn, a[i+1])
			mx = max(mx, a[i+1])
		}
	}

	k := (mn + mx) / 2
	for i := range a {
		if a[i] == -1 {
			a[i] = k
		}
	}

	ans := math.MinInt32
	for i := range n - 1 {
		ans = max(ans, abs(a[i]-a[i+1]))
	}
	fmt.Fprintln(out, ans, k)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
