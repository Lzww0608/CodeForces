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
	for range n - 1 {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	if n&1 == 1 {
		fmt.Fprintln(out, -1)
		return
	}

	ans := 0
	var dfs func(int, int) int
	dfs = func(u, fa int) int {
		sz := 1
		for _, v := range g[u] {
			if v == fa {
				continue
			}

			sz += dfs(v, u)
		}

		if sz&1 == 0 {
			sz = 0
			ans++
		}
		return sz
	}
	dfs(0, -1)

	fmt.Fprintln(out, ans-1)
}
