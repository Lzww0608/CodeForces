package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	b := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] <= k {
			ans = 1
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}

	l, r := -1, -2
	cur := 0
	for i := range n - 1 {
		if b[i]%b[i+1] == 0 {
			if l == -1 {
				l, r = i, i
				cur = a[i]
			}
			cur += a[i+1]
			r++
			for cur > k {
				cur -= a[l]
				l++
			}
			if l >= r {
				l, r = -1, -2
				cur = 0
			}

			ans = max(ans, r-l+1)
		} else {
			l, r = -1, -2
			cur = 0
		}
	}

	fmt.Fprintln(out, ans)
}
