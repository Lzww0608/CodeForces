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

	var s string
	fmt.Fscan(in, &s)
	a := make([]int, 3)
	b := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &b[i])
	}

	var w int
	fmt.Fscan(in, &w)

	cnt := make([]int, 3)
	for i := range s {
		if s[i] == 'B' {
			cnt[0]++
		} else if s[i] == 'S' {
			cnt[1]++
		} else {
			cnt[2]++
		}
	}

	check := func(mid int) bool {
		cost := max(0, mid*cnt[0]-a[0])*b[0] + max(0, mid*cnt[1]-a[1])*b[1] + max(0, mid*cnt[2]-a[2])*b[2]
		return cost <= w
	}

	l, r := -1, 1_000_000_001_000
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
