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
	for i := range n {
		g[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(in, &g[i][j])
		}
	}

	seq := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &seq[i])
	}

	ans := make([]int, n)
	vis := make([]bool, n)
	for k := n - 1; k >= 0; k-- {
		x := seq[k] - 1
		vis[x] = true

		for i := range n {
			for j := range n {
				g[i][j] = min(g[i][j], g[i][x]+g[x][j])
				if vis[i] && vis[j] {
					ans[k] += g[i][j]
				}
			}
		}
	}

	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
}
