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

	var n, m, x, y int
	fmt.Fscan(in, &n, &m)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
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
			fa[min(rx, ry)] = max(rx, ry)
		}
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &x, &y)
		merge(x, y)
	}

	ans := 0
	r := 0
	for l := 1; l <= n; l++ {
		if r > l {
			x = find(l)
			if x < r {
				ans++
				merge(l, r)
			} else {
				r = x
			}
		} else {
			r = max(r, find(l))
		}
	}

	fmt.Fprintln(out, ans)
}
