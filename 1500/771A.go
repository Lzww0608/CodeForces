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
	f := make([]int, n)
	sz := make([]int, n)
	for i := range n {
		f[i] = i
		sz[i] = 1
	}

	var find func(int) int
	find = func(x int) int {
		if x != f[x] {
			f[x] = find(f[x])
		}
		return f[x]
	}

	merge := func(x, y int) {
		rx, ry := find(x), find(y)
		if rx == ry {
			return
		}
		if rx < ry {
			rx, ry = ry, rx
		}
		sz[rx] += sz[ry]
		f[ry] = f[rx]
	}

	var u, v int
	for range m {
		fmt.Fscan(in, &u, &v)
		u--
		v--
		merge(u, v)
	}

	cnt := 0
	for i := range n {
		if i == find(i) {
			cnt += sz[i] * (sz[i] - 1) / 2
		}
	}
	//fmt.Fprintln(out, cnt)
	if cnt == m {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
