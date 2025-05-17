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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	g := make([][]int, n+1)
	var x, y int
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	ans := 0
	var dfs func(int, int, int)
	dfs = func(u, fa, cnt int) {
		if a[u] == 1 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt > m {
			return
		}
		if len(g[u]) == 1 && u != 1 {
			ans++
			return
		}

		for _, v := range g[u] {
			if v != fa {
				dfs(v, u, cnt)
			}
		}
	}
	dfs(1, 0, 0)

	fmt.Fprintln(out, ans)
}
