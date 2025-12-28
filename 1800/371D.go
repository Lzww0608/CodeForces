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
	a := make([]int, n+2)
	sum := make([]int, n+2)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	d := NewDSU(n + 2)

	var q, t, x, y int
	for fmt.Fscan(in, &q); q > 0; q-- {
		fmt.Fscan(in, &t)
		if t == 1 {
			fmt.Fscan(in, &x, &y)
			for y > 0 {
				x = d.find(x)
				if x == n+1 {
					break
				}
				if a[x]-sum[x] >= y {
					sum[x] += y
					y = 0
				} else {
					y -= a[x] - sum[x]
					sum[x] = a[x]
					d.merge(x, x+1)
				}
			}
		} else {
			fmt.Fscan(in, &x)
			fmt.Fprintln(out, sum[x])
		}
	}
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
	y := x
	for d.fa[y] != y {
		y = d.fa[y]
	}
	for i := x; i != y; {
		i, d.fa[i] = d.fa[i], y
	}
	return y
}

func (d *DSU) merge(x, y int) bool {
	rx, ry := d.find(x), d.find(y)
	if rx > ry {
		d.fa[ry] = rx
	} else {
		d.fa[rx] = ry
	}
	return true
}

func (d *DSU) same(x, y int) bool {
	return d.find(x) == d.find(y)
}
