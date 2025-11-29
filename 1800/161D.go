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

	var n, k int
	fmt.Fscan(in, &n, &k)
	g := make([][]int, n+1)
	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	ans := 0
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}

	var dfs func(u, fa int)
	dfs = func(u, fa int) {
		f[u][0] = 1
		for _, v := range g[u] {
			if v == fa {
				continue
			}

			dfs(v, u)
			for i := range k {
				ans += f[u][i] * f[v][k-i-1]
			}
			for i := range k {
				f[u][i+1] += f[v][i]
			}
		}
	}
	dfs(1, 0)
	fmt.Fprintln(out, ans)
}
