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
	h := make([]int, n)
	for i := range h {
		fmt.Fscan(in, &h[i])
	}

	check := func(mid int) bool {
		a, b := 0, 0
		for i := n - 1; i >= 2; i-- {
			x := h[i] + b
			if x < mid {
				return false
			}
			d := min(h[i], x-mid) / 3
			a, b = d*2, d+a
		}

		return h[0]+a >= mid && h[1]+b >= mid
	}

	l, r := 0, h[n-1]+1
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	fmt.Fprintln(out, l)
}
