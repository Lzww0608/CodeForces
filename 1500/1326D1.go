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
	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	base := 0
	for i := 0; i < n/2; i++ {
		if s[i] == s[n-1-i] {
			base++
		} else {
			break
		}
	}

	if n/2 == base {
		fmt.Fprintln(out, s)
		return
	}

	check := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	l, r := base, n-1-base
	for d := 0; d <= r-l; d++ {
		if check(l, r-d) {
			fmt.Fprintln(out, s[:r-d+1]+s[r+1:])
			return
		} else if check(l+d, r) {
			fmt.Fprintln(out, s[:l]+s[l+d:])
			return
		}
	}

	fmt.Fprintln(out, s[:base]+s[n-base:])
}
