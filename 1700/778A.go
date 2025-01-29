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

	var s, t string
	fmt.Fscan(in, &s, &t)
	n := len(s)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	del := make(map[int]bool)

	check := func() bool {
		i, j := 0, 0
		for i < n && j < len(t) {
			if !del[i] && s[i] == t[j] {
				j++
			}
			i++
		}

		return j >= len(t)
	}

	l, r := -1, n+1
	for i := 0; i <= l+((r-l)>>1); i++ {
		del[a[i]-1] = true
	}
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if check() {
			if mid+1 < r {
				for i := mid; i <= mid+((r-mid)>>1); i++ {
					del[a[i]-1] = true
				}
			}
			l = mid
		} else {
			if l+1 < mid {
				for i := mid; i > l+((mid-l)>>1); i-- {
					del[a[i]-1] = false
				}
			}

			r = mid
		}
	}

	fmt.Fprintln(out, l+1)
}
