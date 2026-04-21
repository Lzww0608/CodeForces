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
	a := make([]int, n)
	m := make(map[int][]int)
	for i := range a {
		fmt.Fscan(in, &a[i])
		x := a[i] + i
		if _, ok := m[x]; !ok {
			m[x] = make([]int, 0)
		}
		m[x] = append(m[x], x+i)
	}

	ans := n
	vis := make(map[int]bool)
	var dfs func(int)
	dfs = func(u int) {
		vis[u] = true
		ans = max(ans, u)
		for _, v := range m[u] {
			if !vis[v] {
				dfs(v)
			}
		}
	}
	dfs(n)

	fmt.Fprintln(out, ans)
}
