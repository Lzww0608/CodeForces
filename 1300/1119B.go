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

	var n, h int
	fmt.Fscan(in, &n, &h)
	a := make([]int, n)
	p := make([]int, n)
	for i := range a {
		p[i] = i
		fmt.Fscan(in, &a[i])
	}

	sort.Slice(p, func(i, j int) bool {
		return a[p[i]] > a[p[j]]
	})

	check := func(k int) bool {
		cur, sum := 0, 0
		for _, i := range p {
			if i <= k {
				if cur&1 == 0 {
					sum += a[i]
				}
				cur++
			}
		}

		return sum <= h
	}

	l, r := -1, n
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	fmt.Fprintln(out, l+1)
}
