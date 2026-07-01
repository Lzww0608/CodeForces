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
	g := make([][]pair, n)
	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], pair{v, 0})
		g[v] = append(g[v], pair{u, 1})
	}

	dis := make([]int, n)
	var dfs func(u, fa int) int
	dfs = func(u, fa int) int {
		res := 0
		for _, e := range g[u] {
			if e.u == fa {
				continue
			}
			res += dfs(e.u, u) + e.f
		}
		dis[u] = res
		return res
	}
	dfs(0, -1)

	ans := make([]int, n)
	ans[0] = dis[0]
	minAns := ans[0]

	var dfs2 func(u, fa int)
	dfs2 = func(u, fa int) {
		for _, e := range g[u] {
			if e.u == fa {
				continue
			}
			ans[e.u] = ans[u] + 1 - 2*e.f
			if ans[e.u] < minAns {
				minAns = ans[e.u]
			}
			dfs2(e.u, u)
		}
	}
	dfs2(0, -1)

	fmt.Fprintln(out, minAns)
	for i := 0; i < n; i++ {
		if ans[i] == minAns {
			fmt.Fprint(out, i+1, " ")
		}
	}
}

type pair struct {
	u, f int
}
