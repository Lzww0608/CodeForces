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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	vis := make([]bool, m*n)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	l, r, u, d := 0, 0, 0, 0

	var dfs func(int, int) int
	dfs = func(i, j int) (ans int) {
		if i < 0 || i >= n || j < 0 || j >= m || vis[i*m+j] || s[i][j] == '0' {
			return
		}
		vis[i*m+j] = true
		u = min(u, i)
		d = max(d, i)
		l = min(l, j)
		r = max(l, j)
		ans += 1
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			ans += dfs(x, y)
		}

		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == '1' && !vis[i*m+j] {
				l, r, u, d = j, j, i, i
				cnt := dfs(i, j)
				if cnt != (r-l+1)*(d-u+1) {
					fmt.Fprintln(out, "NO")
					return
				}
			}
		}
	}
	fmt.Fprintln(out, "YES")
	return
}
