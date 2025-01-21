package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x, y int
	fmt.Fscan(in, &n)
	deg := make([]int, n+1)
	g := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &x, &y)
		deg[x]++
		deg[y]++
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	mx := slices.Max(deg) + 1
	col := make([]int, n+1)
	var dfs func(int, int)
	dfs = func(u, fa int) {
		c := 0
		for c == col[u] || c == col[fa] {
			c++
		}
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			col[v] = c
			c = (c + 1) % mx
			dfs(v, u)
			for c == col[u] || c == col[fa] {
				c = (c + 1) % mx
			}
		}
	}
	dfs(1, 0)
	fmt.Fprintln(out, mx)
	for _, x := range col[1:] {
		fmt.Fprint(out, x+1, " ")
	}
}
