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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	fa := make([]int, n)
	mn := make([]int, n)
	for i := range fa {
		fa[i] = i
		mn[i] = a[i]
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
			mn[ry] = min(mn[rx], mn[ry])
		}
	}

	var x, y int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &x, &y)
		x--
		y--
		merge(x, y)
	}

	ans := 0
	vis := make(map[int]bool)
	for i := 0; i < n; i++ {
		t := find(i)
		if !vis[t] {
			vis[t] = true
			ans += mn[t]
		}
	}
	fmt.Fprintln(out, ans)
}
