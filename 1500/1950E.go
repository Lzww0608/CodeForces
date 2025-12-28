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
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	check := func(d int) bool {
		cnt := 0
		for i := range n {
			if s[i] != s[i%d] {
				cnt++
			}
		}
		if cnt <= 1 {
			return true
		}

		cnt = 0
		for i := range n {
			if s[i] != s[n-d+i%d] {
				cnt++
			}
		}
		return cnt <= 1
	}

	ans := n
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		if check(i) {
			fmt.Fprintln(out, i)
			return
		} else if check(n / i) {
			ans = min(ans, n/i)
		}
	}

	fmt.Fprintln(out, ans)
}
