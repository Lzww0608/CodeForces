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
	var a, b, c, n int
	var s string
	fmt.Fscan(in, &n, &a, &b, &c, &s)
	cnt := 0
	aa, bb, cc := a, b, c
	for i := range s {
		if s[i] == 'P' {
			if cc > 0 {
				cnt++
				cc--
			}
		} else if s[i] == 'R' {
			if bb > 0 {
				cnt++
				bb--
			}
		} else {
			if aa > 0 {
				cnt++
				aa--
			}
		}
	}
	if cnt >= (n+1)/2 {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
		return
	}
	ans := make([]byte, n)
	for i := range ans {
		ans[i] = ' '
	}

	for i := range s {
		if s[i] == 'P' {
			if c > 0 {
				ans[i] = 'S'
				c--
			}
		} else if s[i] == 'R' {
			if b > 0 {
				ans[i] = 'P'
				b--
			}
		} else {
			if a > 0 {
				ans[i] = 'R'
				a--
			}
		}
	}
	for i := 0; i < n; i++ {
		if ans[i] == ' ' {
			if c > 0 {
				ans[i] = 'S'
				c--
			} else if b > 0 {
				ans[i] = 'P'
				b--
			} else if a > 0 {
				ans[i] = 'R'
				a--
			}
		}
	}
	fmt.Fprintln(out, string(ans))
	return
}
