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
	var n, m, d int
	fmt.Fscan(in, &n, &m, &d)
	s := make([]int, m)
	for i := range s {
		fmt.Fscan(in, &s[i])
	}

	ans, cnt := 1, 0
	if s[0] == 1 {
		cnt = 1
	}
	for i := range s {
		if s[i] == 1 {
			continue
		}

		ans++
		l := 1
		if i-1 >= 0 {
			l = s[i-1]
		}
		ans += (s[i] - l - 1) / d
	}

	if s[m-1] != n {
		ans += (n - s[m-1]) / d
	}

	res := ans
	for i := range s {
		if s[i] == 1 {
			continue
		}
		l, r := 1, n+1
		if i-1 >= 0 {
			l = s[i-1]
		}
		if i+1 < m {
			r = s[i+1]
		}

		t := res - (s[i]-l-1)/d - (r-s[i]-1)/d + (r-l-1)/d - 1
		if t < ans {
			//fmt.Fprintln(out, t, i)
			ans, cnt = t, 1
		} else if t == ans {
			cnt++
		}
	}

	fmt.Fprintln(out, ans, cnt)
}
