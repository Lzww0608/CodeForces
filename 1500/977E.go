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
	ans := 0
	g := make([][]int, n+1)
	for i := 1; i <= m; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	cnt := 0
	vis := make([]bool, n+1)
	var dfs func(int)
	dfs = func(x int) {
		if vis[x] {
			return
		}
		vis[x] = true
		if len(g[x]) != 2 {
			cnt = 0
		}
		for _, y := range g[x] {
			dfs(y)
		}
	}

	for i := 1; i <= n; i++ {
		if vis[i] {
			continue
		}
		cnt = 1
		dfs(i)
		ans += cnt
	}

	fmt.Fprintln(out, ans)
}
