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
	s := [2]string{}
	for i := 0; i < 2; i++ {
		fmt.Fscan(in, &s[i])
	}
	vis := make([][]bool, 2)
	for i := range vis {
		vis[i] = make([]bool, n)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if vis[i][j] {
			return
		}
		vis[i][j] = true
		for _, v := range [][]int{{i, j - 1}, {i, j + 1}, {i ^ 1, j}} {
			if v[1] >= 0 && v[1] < n {
				if s[v[0]][v[1]] == '<' {
					v[1]--
				} else {
					v[1]++
				}
				dfs(v[0], v[1])
			}
		}
	}
	dfs(0, 0)
	if vis[1][n-1] {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}

}
