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
	var x int
	fmt.Fscan(in, &x)
	n := len(s)

	ans := make([]byte, n)
	for i := range s {
		if s[i] == '0' {
			if i-x >= 0 {
				ans[i-x] = '0'
			}
			if i+x < n {
				ans[i+x] = '0'
			}
		}
	}

	for i := range ans {
		if ans[i] == 0 {
			ans[i] = '1'
		}
	}

	for i := range s {
		if s[i] == '1' {
			if (i-x < 0 || i-x >= 0 && ans[i-x] == '0') && (i+x >= n || i+x < n && ans[i+x] == '0') {
				fmt.Fprintln(out, -1)
				return
			}

		}
	}
	fmt.Fprintln(out, string(ans))
}
