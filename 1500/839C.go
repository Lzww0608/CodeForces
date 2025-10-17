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
	g := make([][]int, n+1)
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	g[1] = append(g[1], -1)

	var ans float64
	var dfs func(int, int, int, float64)
	dfs = func(u, fa, d int, sum float64) {
		if len(g[u]) == 1 {
			ans += sum * float64(d)
			return
		}

		m := 1.0 / float64(len(g[u])-1)
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			dfs(v, u, d+1, sum*m)
		}
	}
	dfs(1, -1, 0, 1.0)
	fmt.Fprintln(out, ans)
}
