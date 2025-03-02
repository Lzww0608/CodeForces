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
	s := make([][]byte, n)
	cnt := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		for j := 0; j < m; j++ {
			if s[i][j] == '#' {
				cnt++
			}
		}
	}
	if cnt <= 2 {
		fmt.Fprintln(out, -1)
		return
	}
	vis := make(map[int]int)

	var dfs func(i, j, t int) int
	dfs = func(i, j, t int) int {
		if i < 0 || i >= n || j < 0 || j >= m || vis[i*m+j] == t || s[i][j] != '#' {
			return 0
		}
		ans := 1
		vis[i*m+j] = t
		ans += dfs(i+1, j, t) + dfs(i, j+1, t) + dfs(i-1, j, t) + dfs(i, j-1, t)
		return ans
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	t := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == '#' {
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x >= 0 && x < n && y >= 0 && y < m && s[x][y] == '#' {
						s[i][j] = '.'
						t++
						if dfs(x, y, t) < cnt-1 {
							fmt.Fprintln(out, 1)
							return
						} else {
							s[i][j] = '#'
							break
						}
					}
				}
			}
		}
	}
	fmt.Fprintln(out, 2)
	return
}
