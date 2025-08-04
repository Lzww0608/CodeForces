package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	g := make([][]int, n)

	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	a := make([]int, n)
	var dfs func(u, fa, d int) int
	dfs = func(u, fa, d int) int {
		sz := 1
		for _, v := range g[u] {
			if v != fa {
				sz += dfs(v, u, d+1)
			}
		}
		a[u] = d + 1 - sz
		return sz
	}
	dfs(0, -1, 0)

	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	ans := 0
	for _, x := range a[:k] {
		ans += x
	}

	fmt.Fprintln(out, ans)
}
