package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	s := make([]string, n)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}

	fa := make([]int, n)
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
			fa[rx] = ry
		}
	}

	vis := make([]int, N)
	for i := range N {
		vis[i] = -1
	}

	for i := range s {
		for j := range s[i] {
			x := int(s[i][j] - 'a')
			if vis[x] == -1 {
				vis[x] = i
			} else {
				merge(vis[x], i)
			}
		}
	}

	ans := 0
	for i := range fa {
		if fa[i] == i {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}
