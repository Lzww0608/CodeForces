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
	s := make([]byte, n)
	fmt.Fscan(in, &s)
	l, r := 0, 0
	for k > 0 {
		for r < n && s[r] == '1' {
			r++
		}

		for l < n && s[l] == '0' {
			l++
		}
		if l == n || r == n {
			break
		}

		if l > r {
			r++
			continue
		}

		d := r - l
		if d < k {
			k -= d
			s[r], s[l] = s[l], s[r]
		} else {
			s[r-k], s[r] = s[r], s[r-k]
			break
		}
	}

	fmt.Fprintln(out, string(s))
}
