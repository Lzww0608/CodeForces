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
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		x--
		y--
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	for i := range n {
		if len(g[i]) != 2 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	for i, v := range g {
		if v == nil {
			continue
		}

		x, pre := 0, -1
		for g[i] != nil {
			x ^= 1
			if g[i][0] != pre {
				i, pre, g[i] = g[i][0], i, nil
			} else {
				i, pre, g[i] = g[i][1], i, nil
			}
		}

		if x != 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}

	fmt.Fprintln(out, "YES")
}
