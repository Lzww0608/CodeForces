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
	var n, m int
	fmt.Fscan(in, &n, &m)
	d := NewDSU(n + m + 1)
	ans := m

	var x, y int
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &x, &y)
		if x == y {
			ans--
			continue
		}

		if !d.merge(i, x+m) {
			ans++
		} else if !d.merge(i, y+m) {
			ans++
		}
	}

	fmt.Fprintln(out, ans)
}

type DSU struct {
	n      int
	fa, sz []int
}

func NewDSU(n int) *DSU {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range n {
		fa[i] = i
		sz[i] = 1
	}
	return &DSU{n: n, fa: fa, sz: sz}
}

func (d *DSU) find(x int) int {
	if d.fa[x] != x {
		d.fa[x] = d.find(d.fa[x])
	}
	return d.fa[x]
}

func (d *DSU) merge(x, y int) bool {
	rx, ry := d.find(x), d.find(y)
	if rx == ry {
		return false
	} else if d.sz[rx] < d.sz[ry] {
		rx, ry = ry, rx
	}
	d.sz[rx] += d.sz[ry]
	d.fa[ry] = rx
	return true
}

func (d *DSU) same(x, y int) bool {
	return d.find(x) == d.find(y)
}
