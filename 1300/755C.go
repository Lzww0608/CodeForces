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
	p := make([]int, n)
	fa := make([]int, n)
	sz := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &p[i])
		fa[i] = i
		sz[i] = 1
		p[i]--
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
		} else if sz[rx] < sz[ry] {
			rx, ry = ry, rx
		}

		sz[rx] += sz[ry]
		fa[ry] = rx
	}

	for i, x := range p {
		merge(i, x)
	}

	ans := 0
	for i, x := range fa {
		if x == i {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}
