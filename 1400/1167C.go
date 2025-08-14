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
	fa := make([]int, n+1)
	sz := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fa[i] = i
		sz[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(x, y int) {
		rx, ry := find(x), find(y)
		if rx != ry {
			fa[rx] = ry
			sz[ry] += sz[rx]
		}
		return
	}

	for i := 0; i < m; i++ {
		pre := -1
		var k, x int
		fmt.Fscan(in, &k)
		for j := 0; j < k; j++ {
			fmt.Fscan(in, &x)
			if pre != -1 {
				merge(pre, x)
			}
			pre = x
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(out, "%d ", sz[find(i)])
	}
}
