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
	var n int
	var s string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)
	cnt := 0
	for i := range s {
		if s[i] == '0' {
			cnt++
		} else {
			cnt--
		}
	}
	if cnt != 0 {
		fmt.Fprintln(out, -1)
		return
	}
	ans := []int{}
	for {
		l, r := 0, len(s)-1
		for r >= 0 && s[l] != s[r] {
			l++
			r--
		}

		if r < l {
			break
		}

		if s[l] == s[r] {
			if s[l] == '1' {
				ans = append(ans, l)
				s = s[:l] + "01" + s[l:]
			} else {
				ans = append(ans, r+1)
				s = s[:r+1] + "01" + s[r+1:]
			}
		}

	}
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
