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

	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	cnt := 0
	f := make([][]byte, n)
	for i := range f {
		fmt.Fscan(in, &f[i])
		for j := range f[i] {
			if f[i][j] == '.' {
				f[i][j] = 'X'
				cnt++
			}
		}
	}

	k = cnt - k

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || k <= 0 || f[i][j] != 'X' {
			return
		}
		k--
		f[i][j] = '.'
		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i-1, j)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if f[i][j] == 'X' {
				dfs(i, j)
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintln(out, string(f[i]))
	}
}
