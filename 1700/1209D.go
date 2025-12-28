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

	var n, k int
	fmt.Fscan(in, &n, &k)
	var u, v int
	d := NewDSU(n + 1)
	for range k {
		fmt.Fscan(in, &u, &v)
		d.merge(u, v)
	}

	ans := k
	for i := 1; i <= n; i++ {
		if d.fa[i] == i {
			x := d.sz[i]
			ans -= x - 1
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
