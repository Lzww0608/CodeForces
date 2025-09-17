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
	var n int
	fmt.Fscan(in, &n)
	g := make([][]int, n+1)
	for i := 2; i <= n; i++ {
		var j int
		fmt.Fscan(in, &j)
		g[j] = append(g[j], i)
		g[i] = append(g[i], j)
	}

	var s string
	fmt.Fscan(in, &s)

	ans := 0

	var dfs func(int, int) int
	dfs = func(u, fa int) (cnt int) {
		for _, v := range g[u] {
			if v == fa {
				continue
			}
			cnt += dfs(v, u)
		}

		if s[u-1] == 'W' {
			cnt++
		} else {
			cnt--
		}

		if cnt == 0 {
			ans++
		}

		return
	}
	dfs(1, 0)

	fmt.Fprintln(out, ans)
}
