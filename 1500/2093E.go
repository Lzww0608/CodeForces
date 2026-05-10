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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	check := func(mid int) bool {
		cnt := 0
		vis := make([]bool, mid+1)
		t := 0
		for _, x := range a {
			if x >= t && x <= mid {
				vis[x] = true
			}

			for vis[t] && t < mid {
				vis[t] = false
				t++
			}

			if t >= mid {
				t = 0
				cnt++
			}
		}

		return cnt >= k
	}

	l, r := 0, n+1
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
