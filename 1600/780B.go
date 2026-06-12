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
	x := make([]int, n)
	v := make([]int, n)
	for i := range x {
		fmt.Fscan(in, &x[i])
	}

	for i := range v {
		fmt.Fscan(in, &v[i])
	}

	check := func(mid float64) bool {
		l, r := -1e15, 1e15
		for i := range n {
			l = max(l, float64(x[i])-float64(v[i])*mid)
			r = min(r, float64(x[i])+float64(v[i])*mid)
			if l+1e-7 > r {
				return false
			}
		}
		return true
	}

	l, r := 0.0, 1000000000.0
	for l+1e-7 < r {
		mid := l + ((r - l) / 2)
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}

	fmt.Fprintln(out, r)
}
