package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k, q int
	fmt.Fscan(in, &n, &k, &q)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, k)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
		}
	}

	for j := 0; j < k; j++ {
		for i := 1; i < n; i++ {
			a[i][j] |= a[i-1][j]
		}
	}
	//fmt.Fprintln(out, a)

	var x, y, t int
	var s string
	for ; q > 0; q-- {
		l, r := 0, n-1
		for fmt.Fscan(in, &t); t > 0; t-- {
			fmt.Fscan(in, &x, &s, &y)
			//fmt.Fprintln(out, x, s, y)
			x--
			c := s[0]
			if c == '>' {
				pos := sort.Search(n, func(i int) bool {
					return a[i][x] > y
				})
				l = max(l, pos)
			} else {
				pos := sort.Search(n, func(i int) bool {
					return a[i][x] >= y
				})
				r = min(r, pos-1)
			}
		}
		if l <= r {
			fmt.Fprintln(out, l+1)
		} else {
			fmt.Fprintln(out, -1)
		}
	}
}
