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
	a := make([][]int, n)

	b := true
	for i := range a {
		var k int
		fmt.Fscan(in, &k)
		if k != 0 {
			b = false
		}
		a[i] = make([]int, k)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			a[i][j]--
		}
	}
	if b {
		fmt.Fprintln(out, n)
		return
	}

	f := make([]int, n)
	for i := range f {
		f[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if f[x] != x {
			f[x] = find(f[x])
		}
		return f[x]
	}

	merge := func(x, y int) {
		rx, ry := find(x), find(y)
		if rx != ry {
			f[rx] = ry
		}
	}

	for i := 0; i < n; i++ {
	next:
		for j := i + 1; j < n; j++ {
			if find(i) == find(j) {
				continue
			}
			for _, x := range a[i] {
				for _, y := range a[j] {
					if x == y {
						merge(i, j)
						continue next
					}
				}
			}
		}
	}

	cnt := 0
	for i, x := range f {
		if x == i {
			cnt++
		}
	}

	fmt.Fprintln(out, cnt-1)
}
