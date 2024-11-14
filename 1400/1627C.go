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
	var n, x, y int
	fmt.Fscan(in, &n)
	g := make([][]int, n)
	ans := make([]int, n-1)
	m := map[int]map[int]int{}
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &x, &y)
		x--
		y--
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		if len(m[x]) == 0 {
			m[x] = map[int]int{}
		}
		if len(m[y]) == 0 {
			m[y] = map[int]int{}
		}
		m[x][y] = i
		m[y][x] = i
	}
	for i := range g {
		if len(g[i]) > 2 {
			fmt.Fprintln(out, -1)
			return
		}
	}

	var dfs func(int, int, int)
	dfs = func(x, fa, pre int) {
		for _, y := range g[x] {
			if y == fa {
				continue
			}
			ans[m[x][y]] = pre + 2
			pre ^= 1
			dfs(y, x, pre)
		}
	}
	dfs(0, -1, 0)
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
