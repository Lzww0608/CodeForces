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
	var n, m1, m2 int
	fmt.Fscan(in, &n, &m1, &m2)
	d1 := NewDSU(n + 1)
	d2 := NewDSU(n + 1)

	var u, v int
	g := make([][2]int, m1)
	for i := range m1 {
		fmt.Fscan(in, &u, &v)
		g[i] = [2]int{u, v}
	}
	f := make([][2]int, m2)
	for i := range m2 {
		fmt.Fscan(in, &u, &v)
		d2.merge(u, v)
		f[i] = [2]int{u, v}
	}

	ans := 0
	for _, e := range g {
		u, v = e[0], e[1]
		if d2.same(u, v) {
			d1.merge(u, v)
		} else {
			ans++
		}
	}

	for _, e := range f {
		u, v = e[0], e[1]
		if !d1.same(u, v) {
			ans++
			d1.merge(u, v)
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

func (d *DSU) merge(x, y int) {
	rx, ry := d.find(x), d.find(y)
	if rx == ry {
		return
	} else if d.sz[rx] < d.sz[ry] {
		rx, ry = ry, rx
	}
	d.sz[rx] += d.sz[ry]
	d.fa[ry] = rx
	return
}

func (d *DSU) same(x, y int) bool {
	return d.find(x) == d.find(y)
}
