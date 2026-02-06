package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	g := make([][]int, n)

	for i := 1; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		x--
		g[i] = append(g[i], x)
		g[x] = append(g[x], i)
	}

	check := func(mid int) bool {
		var dfs func(int, int, int) bool
		dfs = func(u, fa, t int) bool {
			if t > int(1e9) {
				return false
			}
			cur := t
			if a[u] < t {
				if len(g[u]) == 1 {
					return false
				}
				cur += t - a[u]
			}
			for _, v := range g[u] {
				if v == fa {
					continue
				}
				if !dfs(v, u, cur) {
					return false
				}
			}

			return true
		}

		for _, v := range g[0] {
			if !dfs(v, 0, mid-a[0]) {
				return false
			}
		}
		return true
	}

	l, r := a[0], int(2e9)+1
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	fmt.Fprintln(out, l)
}
