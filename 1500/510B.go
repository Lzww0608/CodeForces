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
	s := make([]string, n)
	vis := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		vis[i] = make([]int, m)
	}

	ans := false

	var dfs func(int, int, int, byte)
	dfs = func(x, y, c int, ch byte) {
		if ans || x < 0 || x >= n || y < 0 || y >= m || s[x][y] != ch {
			return
		}
		if vis[x][y] != 0 {
			if c-vis[x][y] >= 4 {
				ans = true
			}
			return
		}
		vis[x][y] = c
		dfs(x+1, y, c+1, ch)
		dfs(x-1, y, c+1, ch)
		dfs(x, y+1, c+1, ch)
		dfs(x, y-1, c+1, ch)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vis[i][j] == 0 {
				dfs(i, j, 1, s[i][j])
				if ans {
					fmt.Fprintln(out, "Yes")
					return
				}
			}
		}
	}

	fmt.Fprintln(out, "No")
}
