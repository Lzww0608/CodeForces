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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)

	fa := make([]int, n)
	sz := make([]int, n)
	cnt := make([][N]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = 1
		x := int(s[i] - 'a')
		cnt[i][x] = 1
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
		for i := range N {
			cnt[rx][i] += cnt[ry][i]
		}
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		merge(i, j)
	}

	for i := k; i < n; i++ {
		merge(i, i-k)
	}

	ans := 0
	for i := range n {
		x := find(i)
		if x != i {
			continue
		}

		mx := 0
		for j := range N {
			mx = max(mx, cnt[x][j])
		}

		ans += sz[x] - mx
	}

	fmt.Fprintln(out, ans)
}
