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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, c, d int
	fmt.Fscan(in, &n, &c, &d)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + a[i]
	}

	if a[0]*d < c {
		fmt.Fprintln(out, "Impossible")
		return
	}

	if pre[min(n, d)] >= c {
		fmt.Fprintln(out, "Infinity")
		return
	}

	can := func(k int) bool {
		cycle := k + 1
		full := d / cycle
		rem := d % cycle

		useInCycle := min(n, cycle)
		useInRem := min(n, rem)

		total := full*pre[useInCycle] + pre[useInRem]
		return total >= c
	}

	l, r := -1, d
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if can(mid) {
			l = mid
		} else {
			r = mid
		}
	}

	fmt.Fprintln(out, l)
}
