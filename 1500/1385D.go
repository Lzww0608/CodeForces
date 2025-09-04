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
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	var dfs func(l, r, c int) int
	dfs = func(l, r, c int) int {
		if l == r {
			if int(s[l]-'a') == c {
				return 0
			}
			return 1
		}

		cnt := 0
		var i int
		for i = l; i <= (r+l)/2; i++ {
			if int(s[i]-'a') != c {
				cnt++
			}
		}
		ans := cnt + dfs(i, r, c+1)

		cnt = dfs(l, i-1, c+1)
		for ; i <= r; i++ {
			if int(s[i]-'a') != c {
				cnt++
			}
		}
		ans = min(ans, cnt)
		return ans
	}

	fmt.Fprintln(out, dfs(0, n-1, 0))
}
