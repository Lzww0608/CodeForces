package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/356/A
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	fa := make([]int, n+2)
	ans := make([]int, n+1)
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

	var l, r, x int
	for range m {
		fmt.Fscan(in, &l, &r, &x)
		for l = find(l); l <= r; l = find(l + 1) {
			if l != x {
				ans[l] = x
				fa[l] = l + 1
			}
		}
	}

	for _, x := range ans[1:] {
		fmt.Fprintf(out, "%d ", x)
	}
}
