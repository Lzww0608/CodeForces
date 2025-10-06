package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const MOD int = 998244353

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

	f := make(map[int]int)
	f[0] = 1
	ans := 0
	for _, x := range a {
		g := make(map[int]int)
		mn, mx := math.MaxInt, math.MinInt
		for k, v := range f {
			g[k+x] += v
			g[abs(k+x)] += v
			mn = min(mn, k+x)
			mx = max(mx, abs(k+x))
		}

		clear(f)
		f[mn] = g[mn] % MOD
		f[mx] = g[mx] % MOD
		ans = f[mx]
	}

	fmt.Fprintln(out, ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
