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
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	b := make([]int, m)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	check := func(mid int) bool {
		i, j := 0, 1
		l, r := b[0]-mid, b[0]+mid
		for ; i < n; i++ {
			if a[i] < l {
				return false
			} else if a[i] <= r {
				continue
			}
			for j < m && a[i] > r {
				l, r = b[j]-mid, b[j]+mid
				j++
			}

			if a[i] > r || a[i] < l {
				return false
			}
		}

		return true
	}

	l, r := -1, int(2e9)+1
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
