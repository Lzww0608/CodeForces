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
	a := make([]int, n)
	d := NewDSU(n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
		d.merge(i, a[i])
	}

	cycle := 0
	b := make([]int, n)
	for i := range n {
		if a[a[i]] == i {
			cycle = 1
			b[d.find(i)] = 1
		}

	}

	mn, mx := cycle, 0
	for i := range n {
		if j := d.find(i); j == i {
			mx++
			if b[j] == 0 {
				mn++
			}
		}
	}

	fmt.Fprintln(out, mn, mx)
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
	if d.sz[rx] < d.sz[ry] {
		rx, ry = ry, rx
	}
	d.sz[rx] += d.sz[ry]
	d.fa[ry] = d.fa[rx]
}
