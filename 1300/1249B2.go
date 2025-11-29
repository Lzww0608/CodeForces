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
	fa := make([]int, n)
	sz := make([]int, n)
	a := make([]int, n)
	for i := range a {
		fa[i] = i
		sz[i] = 1
		fmt.Fscan(in, &a[i])
		a[i]--
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
		if rx == ry {
			return
		}

		if sz[rx] < sz[ry] {
			rx, ry = ry, rx
		}
		fa[ry] = rx
		sz[rx] += sz[ry]
	}

	for i, x := range a {
		merge(i, x)
	}

	for i := range n {
		fmt.Fprintf(out, "%d ", sz[find(i)])
	}

	fmt.Fprintln(out)
}
