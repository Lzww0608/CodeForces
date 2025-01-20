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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, x, y int
	fmt.Fscan(in, &n)
	g := make([][]int, n+1)
	sz := make([]int, n+1)
	mx := make([]int, n+1)
	ct := []int{}
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &x, &y)
		g[y] = append(g[y], x)
		g[x] = append(g[x], y)
	}

	var dfs func(u, fa int)
	dfs = func(u, fa int) {
		sz[u] = 1
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			dfs(v, u)
			sz[u] += sz[v]
			mx[u] = max(mx[u], sz[v])
		}
		mx[u] = max(mx[u], n-sz[u])
		if mx[u] <= n/2 {
			ct = append(ct, u)
		}
	}

	if n&1 == 1 {
		fmt.Fprintln(out, g[1][0], 1)
		fmt.Fprintln(out, 1, g[1][0])
	} else {
		dfs(1, 0)
		if len(ct) == 1 {
			fmt.Fprintln(out, 1, g[1][0])
			fmt.Fprintln(out, g[1][0], 1)
		} else {
			a, b := ct[0], ct[1]
			h := 1
			for _, v := range g[a] {
				if v != b {
					h = v
					break
				}
			}
			fmt.Fprintln(out, a, h)
			fmt.Fprintln(out, h, b)
		}
	}

}
