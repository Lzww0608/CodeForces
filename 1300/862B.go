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

	var n int
	fmt.Fscan(in, &n)
	g := make([][]int, n)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	color := make([]int, n)
	var dfs func(int, int, int)
	dfs = func(u, fa, c int) {
		color[u] = c
		for _, v := range g[u] {
			if v != fa {
				dfs(v, u, c^3)
			}
		}
	}
	dfs(0, -1, 1)
	cnt := 0
	for _, x := range color {
		if x == 1 {
			cnt++
		}
	}

	fmt.Fprintln(out, cnt*(n-cnt)-n+1)
}
