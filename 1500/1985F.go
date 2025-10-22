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
	var h, n int
	fmt.Fscan(in, &h, &n)
	a := make([]int, n)
	c := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		sum += a[i]
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &c[i])
	}

	check := func(mid int) bool {
		cur := sum
		if cur >= h {
			return true
		}
		mid -= 1
		for i, x := range c {
			cur += (mid / x) * a[i]
		}
		return cur >= h
	}

	l, r := 0, int(1e11)
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Fprintln(out, r)
}
