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
	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]int, n)
	for range m {
		var u, v int
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	cnt := [3]int{}
	col := make([]int, n)
	var dfs func(int, int)
	dfs = func(u, c int) {
		if col[u] != 0 {
			return
		}
		col[u] = c
		cnt[c]++
		for _, v := range g[u] {
			dfs(v, 3-c)
		}
	}

	dfs(0, 1)

	choose := 1
	if cnt[1] > cnt[2] {
		choose = 2
	}

	fmt.Fprintln(out, cnt[choose])
	for i := 0; i < n; i++ {
		if col[i] == choose {
			fmt.Fprint(out, i+1, " ")
		}
	}

	fmt.Fprintln(out)
}
